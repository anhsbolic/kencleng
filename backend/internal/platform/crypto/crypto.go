package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// gcmNonceSize is the standard 12-byte nonce length used by AES-GCM.
const gcmNonceSize = 12

// gcmTagSize is the 16-byte authentication tag appended by GCM. Decrypt
// requires at least nonce + tag bytes to have any ciphertext at all.
const gcmTagSize = 16

// Encrypt encrypts plaintext using AES-GCM with a random 12-byte nonce
// derived from crypto/rand. The returned ciphertext is
// nonce || ciphertext || tag, suitable for storage in a BYTEA column.
// The same plaintext encrypts to a different ciphertext on every call
// (random nonce), so equality checks must use HMAC, not ciphertext
// comparison.
//
// keys.EncryptionKey must be 32 bytes (AES-256); this matches the
// validation already enforced by New when the Keys holder is
// constructed.
func Encrypt(plaintext []byte, keys *Keys) ([]byte, error) {
	aead, err := newAEAD(keys)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcmNonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("crypto: encrypt: read nonce: %w", err)
	}
	// Seal appends the encrypted plaintext and the authentication tag to
	// dst (the nonce), producing nonce || ciphertext || tag in one call.
	return aead.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts a ciphertext produced by Encrypt. The expected layout
// is nonce (first 12 bytes) || ciphertext || tag. A non-nil error is
// returned (without leaking internals beyond the category) if the input
// is too short or the authentication tag does not verify — the latter
// indicates tampering or a wrong key.
func Decrypt(ciphertext []byte, keys *Keys) ([]byte, error) {
	aead, err := newAEAD(keys)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < gcmNonceSize+gcmTagSize {
		return nil, fmt.Errorf("crypto: decrypt: ciphertext too short (got %d bytes)", len(ciphertext))
	}
	nonce := ciphertext[:gcmNonceSize]
	ct := ciphertext[gcmNonceSize:]
	plaintext, err := aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, fmt.Errorf("crypto: decrypt: %w", err)
	}
	return plaintext, nil
}

// HMAC computes HMAC-SHA256 over data using keys.HMACKey and returns the
// lowercase hex digest. The output is deterministic for a fixed
// (data, key) pair: the same plaintext always produces the same digest,
// which is what makes the *_hash columns usable for uniqueness and
// lookups. keys.HMACKey must be 32 bytes (validated by New).
func HMAC(data []byte, keys *Keys) string {
	mac := hmac.New(sha256.New, keys.HMACKey)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// newAEAD builds an AES-GCM AEAD from keys.EncryptionKey. It is cheap to
// construct and is built per call rather than cached on Keys so that
// Keys stays a plain, immutable data holder with no methods of its own.
func newAEAD(keys *Keys) (cipher.AEAD, error) {
	if keys == nil {
		return nil, fmt.Errorf("crypto: keys are nil")
	}
	if len(keys.EncryptionKey) != 32 {
		return nil, fmt.Errorf("crypto: encryption key must be 32 bytes, got %d", len(keys.EncryptionKey))
	}
	block, err := aes.NewCipher(keys.EncryptionKey)
	if err != nil {
		return nil, fmt.Errorf("crypto: new aes cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("crypto: new gcm: %w", err)
	}
	return aead, nil
}
