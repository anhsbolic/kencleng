package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// testKeys returns a fixed, non-secret Keys pair (32-byte zero keys) for
// deterministic tests. Real keys come from base64 env vars via New; tests
// never touch those.
func testKeys(t *testing.T) *Keys {
	t.Helper()
	return &Keys{
		EncryptionKey: bytes.Repeat([]byte{0x01}, 32),
		HMACKey:       bytes.Repeat([]byte{0x02}, 32),
	}
}

// TestEncryptDecrypt_RoundTrip verifies Decrypt recovers the exact
// plaintext that Encrypt consumed.
func TestEncryptDecrypt_RoundTrip(t *testing.T) {
	keys := testKeys(t)
	plaintext := []byte("user@example.com")

	ct, err := Encrypt(plaintext, keys)
	if err != nil {
		t.Fatalf("Encrypt: unexpected error: %v", err)
	}
	pt, err := Decrypt(ct, keys)
	if err != nil {
		t.Fatalf("Decrypt: unexpected error: %v", err)
	}
	if !bytes.Equal(pt, plaintext) {
		t.Fatalf("round-trip mismatch: got %q, want %q", pt, plaintext)
	}
}

// TestEncryptDecrypt_RoundTrip_EmptyPlaintext verifies the empty-string
// case still round-trips (valid plaintext, produces nonce+tag only).
func TestEncryptDecrypt_RoundTrip_EmptyPlaintext(t *testing.T) {
	keys := testKeys(t)
	ct, err := Encrypt(nil, keys)
	if err != nil {
		t.Fatalf("Encrypt(nil): %v", err)
	}
	pt, err := Decrypt(ct, keys)
	if err != nil {
		t.Fatalf("Decrypt: %v", err)
	}
	if len(pt) != 0 {
		t.Fatalf("expected empty plaintext, got %q", pt)
	}
}

// TestEncrypt_NonceUniqueness verifies two encryptions of the same
// plaintext yield different ciphertexts (random nonce), so ciphertext
// comparison must never be used for equality.
func TestEncrypt_NonceUniqueness(t *testing.T) {
	keys := testKeys(t)
	plaintext := []byte("same@example.com")

	a, err := Encrypt(plaintext, keys)
	if err != nil {
		t.Fatalf("Encrypt (a): %v", err)
	}
	b, err := Encrypt(plaintext, keys)
	if err != nil {
		t.Fatalf("Encrypt (b): %v", err)
	}
	if bytes.Equal(a, b) {
		t.Fatal("two encryptions of the same plaintext produced identical ciphertext (nonce not random)")
	}
}

// TestDecrypt_TamperRejected verifies flipping any byte of the ciphertext
// causes Decrypt to fail (GCM authentication), so tampering is detected.
func TestDecrypt_TamperRejected(t *testing.T) {
	keys := testKeys(t)
	ct, err := Encrypt([]byte("sensitive@example.com"), keys)
	if err != nil {
		t.Fatalf("Encrypt: %v", err)
	}

	for i := range ct {
		tampered := make([]byte, len(ct))
		copy(tampered, ct)
		tampered[i] ^= 0xFF
		if _, err := Decrypt(tampered, keys); err == nil {
			t.Fatalf("Decrypt succeeded after flipping byte %d (tamper not detected)", i)
		}
	}
}

// TestDecrypt_WrongKeyRejected verifies decryption with a different key
// fails (auth tag does not verify).
func TestDecrypt_WrongKeyRejected(t *testing.T) {
	keys := testKeys(t)
	ct, err := Encrypt([]byte("secret@example.com"), keys)
	if err != nil {
		t.Fatalf("Encrypt: %v", err)
	}
	wrong := &Keys{
		EncryptionKey: bytes.Repeat([]byte{0x03}, 32),
		HMACKey:       bytes.Repeat([]byte{0x04}, 32),
	}
	if _, err := Decrypt(ct, wrong); err == nil {
		t.Fatal("Decrypt succeeded with a different key (auth not enforced)")
	}
}

// TestDecrypt_ShortCiphertextErrors verifies inputs shorter than
// nonce+tag are rejected without panicking.
func TestDecrypt_ShortCiphertextErrors(t *testing.T) {
	keys := testKeys(t)
	short := []byte("too short")
	if _, err := Decrypt(short, keys); err == nil {
		t.Fatal("Decrypt succeeded on a too-short ciphertext")
	}
	// Exactly the minimum boundary: nonce + tag, no ciphertext — a valid
	// empty-plaintext ciphertext. Decrypt of a byte short of that must error.
	tooShortByOne := make([]byte, gcmNonceSize+gcmTagSize-1)
	if _, err := Decrypt(tooShortByOne, keys); err == nil {
		t.Fatal("Decrypt succeeded on ciphertext shorter than nonce+tag")
	}
}

// TestHMAC_Determinism verifies the same (data, key) pair always yields
// the same digest — the property the *_hash lookaside columns rely on
// for uniqueness and lookups.
func TestHMAC_Determinism(t *testing.T) {
	keys := testKeys(t)
	data := []byte("user@example.com")

	first := HMAC(data, keys)
	for i := 0; i < 5; i++ {
		if got := HMAC(data, keys); got != first {
			t.Fatalf("HMAC not deterministic: got %q, want %q", got, first)
		}
	}
}

// TestHMAC_KeyDependence verifies a different key yields a different
// digest.
func TestHMAC_KeyDependence(t *testing.T) {
	data := []byte("user@example.com")
	a := HMAC(data, testKeys(t))
	keysB := &Keys{
		EncryptionKey: bytes.Repeat([]byte{0x01}, 32),
		HMACKey:       bytes.Repeat([]byte{0x09}, 32), // different HMAC key
	}
	b := HMAC(data, keysB)
	if a == b {
		t.Fatal("HMAC with a different key produced the same digest")
	}
}

// TestHMAC_Format verifies the output is a 64-char lowercase hex string
// (SHA-256 = 32 bytes = 64 hex chars).
func TestHMAC_Format(t *testing.T) {
	keys := testKeys(t)
	got := HMAC([]byte("x"), keys)
	if len(got) != 64 {
		t.Fatalf("HMAC output length = %d, want 64", len(got))
	}
	if _, err := hex.DecodeString(got); err != nil {
		t.Fatalf("HMAC output is not valid lowercase hex: %v (got %q)", err, got)
	}
}

// TestEncrypt_NilKeysErrors verifies Encrypt fails fast on a nil Keys
// rather than panicking.
func TestEncrypt_NilKeysErrors(t *testing.T) {
	if _, err := Encrypt([]byte("x"), nil); err == nil {
		t.Fatal("Encrypt(nil keys) should error")
	}
}

// TestDecrypt_NilKeysErrors verifies Decrypt fails fast on a nil Keys.
func TestDecrypt_NilKeysErrors(t *testing.T) {
	if _, err := Decrypt([]byte("x"), nil); err == nil {
		t.Fatal("Decrypt(nil keys) should error")
	}
}

// TestEncrypt_BadKeySizeErrors verifies Encrypt rejects a wrong-size
// encryption key (defense in depth — New already enforces 32 bytes).
func TestEncrypt_BadKeySizeErrors(t *testing.T) {
	bad := &Keys{
		EncryptionKey: bytes.Repeat([]byte{0x01}, 16), // AES-128, not 256
		HMACKey:       bytes.Repeat([]byte{0x02}, 32),
	}
	if _, err := Encrypt([]byte("x"), bad); err == nil {
		t.Fatal("Encrypt with a 16-byte key should error (must be 32)")
	}
}
