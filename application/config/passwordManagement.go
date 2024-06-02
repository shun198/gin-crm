package config

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(result), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomPassword() (string, error) {
	token, err := TokenGenerator(32)
	if err != nil {
		return "", err
	}
	return token, nil
}
