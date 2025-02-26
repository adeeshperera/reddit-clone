package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

const defaultTokenExpiration = time.Hour * 24

// JWTManager handles the generation and verification of JWT tokens
type JWTManager struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

// NewJWTManager initializes a new JWTManager with keys loaded from PEM files
func NewJWTManager(privateKeyPath, publicKeyPath string) (*JWTManager, error) {
	privateKey, err := loadECDSAKey(privateKeyPath, true)
	if err != nil {
		return nil, err
	}

	publicKey, err := loadECDSAKey(publicKeyPath, false)
	if err != nil {
		return nil, err
	}

	return &JWTManager{
		privateKey: privateKey.(*ecdsa.PrivateKey),
		publicKey:  publicKey.(*ecdsa.PublicKey),
	}, nil
}

// GenerateToken generates a new JWT token for a user
func (jm *JWTManager) GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(defaultTokenExpiration).Unix(),
		"iat":     time.Now().Unix(), // Issued at
		"nbf":     time.Now().Unix(), // Not before
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	return token.SignedString(jm.privateKey)
}

// VerifyToken verifies and parses a JWT token
func (jm *JWTManager) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jm.publicKey, nil
	})
}

// loadECDSAKey loads an ECDSA private or public key from a PEM file
func loadECDSAKey(path string, isPrivate bool) (interface{}, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading key file: %v", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	if isPrivate {
		return x509.ParseECPrivateKey(block.Bytes)
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	return pubKey.(*ecdsa.PublicKey), nil
}
