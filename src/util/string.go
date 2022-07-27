package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashSecret(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	return string(bytes), err
}

func CompareHashAndSecret(hash string, secret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
