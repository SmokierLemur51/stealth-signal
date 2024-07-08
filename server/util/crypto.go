package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashMessage(s string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashedByte), err
}

func HashMessageNoError(s string) string {
	hashedByte, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashedByte)
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
