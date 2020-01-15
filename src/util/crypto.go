package util

import (
	"crypto/rand"
	"crypto/rsa"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2"
)

// Crypto s
type Crypto struct{}

// COST ss
const COST int = 7

// Hash s
func Hash(password []byte) []byte {
	if password == nil {
		return nil
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), COST)
		if err != nil {
			return nil
		}
		return hash
	}
}

// Compare s
func Compare(hash []byte, password []byte) bool {
	if hash == nil || password == nil {
		return false
	} else {
		err := bcrypt.CompareHashAndPassword(hash, password)
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

// CreateEncryptedJWT s
func CreateEncryptedJWT(payload string) (*jose.JSONWebEncryption, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: publicKey}, nil)
	if err != nil {
		return nil, nil, err
	}

	var plaintext = []byte(payload)
	object, err := encrypter.Encrypt(plaintext)
	if err != nil {
		return nil, nil, err
	}

	encryption := object

	return encryption, privateKey, nil
}

// VerifyEncryptedToken s
func VerifyEncryptedToken(payload string) (*jose.JSONWebEncryption, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.RSA_OAEP_256, Key: publicKey}, nil)
	if err != nil {
		return nil, nil, err
	}

	var plaintext = []byte(payload)
	object, err := encrypter.Encrypt(plaintext)
	if err != nil {
		return nil, nil, err
	}

	encryption := object

	return encryption, privateKey, nil
}
