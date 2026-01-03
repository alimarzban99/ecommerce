package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// LoadKeys loads RSA private and public keys from file paths
func LoadKeys(privateKeyPath, publicKeyPath string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// Load private key
	privateKeyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse private key: %w", err)
	}

	// Load public key
	publicKeyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("could not read public key: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse public key: %w", err)
	}

	return privateKey, publicKey, nil
}
