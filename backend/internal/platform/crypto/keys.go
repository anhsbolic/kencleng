package crypto

import (
	"encoding/base64"
	"fmt"
)

// Keys holds the raw cryptographic key material loaded from environment
// configuration. EncryptionKey is the AES-GCM key; HMACKey is the key used
// for the deterministic *_hash lookaside columns.
type Keys struct {
	EncryptionKey []byte
	HMACKey       []byte
}

// New parses and validates the base64-encoded encryption and HMAC keys. Both
// keys must be non-empty and decode to exactly 32 bytes, matching the format
// documented in docs/project/kencleng-backend-tech-stack.md.
func New(encryptionKeyB64, hmacKeyB64 string) (*Keys, error) {
	encryptionKey, err := decodeKey("ENCRYPTION_KEY", encryptionKeyB64)
	if err != nil {
		return nil, err
	}
	hmacKey, err := decodeKey("HMAC_KEY", hmacKeyB64)
	if err != nil {
		return nil, err
	}
	return &Keys{EncryptionKey: encryptionKey, HMACKey: hmacKey}, nil
}

func decodeKey(name, b64 string) ([]byte, error) {
	if b64 == "" {
		return nil, fmt.Errorf("%s is empty", name)
	}
	key, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("%s is not valid base64: %w", name, err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("%s must decode to 32 bytes, got %d", name, len(key))
	}
	return key, nil
}
