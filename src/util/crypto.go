package util

import "golang.org/x/crypto/bcrypt"

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

// Valid s
func Valid(hash []byte, password []byte) bool {
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
