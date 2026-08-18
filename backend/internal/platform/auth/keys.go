package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Keys holds the ES256 signing/verification key pair loaded at startup. The
// access token is signed with the private key and verified with the public
// key; see docs/project/kencleng-backend-tech-stack.md.
type Keys struct {
	Private *ecdsa.PrivateKey
	Public  *ecdsa.PublicKey
}

// Load reads the PEM-encoded EC private and public keys from the given paths
// and verifies they form a consistent pair before returning them.
func Load(privatePath, publicPath string) (*Keys, error) {
	privateKey, err := loadPrivateKey(privatePath)
	if err != nil {
		return nil, err
	}
	publicKey, err := loadPublicKey(publicPath)
	if err != nil {
		return nil, err
	}
	if !privateKey.PublicKey.Equal(publicKey) {
		return nil, fmt.Errorf("public key %s does not match private key %s", publicPath, privatePath)
	}
	return &Keys{Private: privateKey, Public: publicKey}, nil
}

func loadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read private key: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("private key %s: no PEM block found", path)
	}
	// SEC1 "EC PRIVATE KEY" (openssl ecparam/ec output).
	if key, err := x509.ParseECPrivateKey(block.Bytes); err == nil {
		return key, nil
	}
	// PKCS#8 "PRIVATE KEY" fallback.
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		ec, ok := key.(*ecdsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("private key %s is not an EC key", path)
		}
		return ec, nil
	}
	return nil, fmt.Errorf("private key %s: unsupported PEM block %q", path, block.Type)
}

func loadPublicKey(path string) (*ecdsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read public key: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("public key %s: no PEM block found", path)
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("public key %s: %w", path, err)
	}
	ec, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key %s is not an EC key", path)
	}
	return ec, nil
}
