package config

import (
	"errors"
	"regexp"

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

func ValidatePassword(password string) error {
	upperCase := regexp.MustCompile(`[A-Z]`)
	lowerCase := regexp.MustCompile(`[a-z]`)
	numbers := regexp.MustCompile(`[0-9]`)
	passwordLength := len(password)

	if match := upperCase.MatchString(password); !match {
		return errors.New("パスワードに英字大文字を入れてください")
	} else if match := lowerCase.MatchString(password); !match {
		return errors.New("パスワードに英字小文字を入れてください")
	} else if match := numbers.MatchString(password); !match {
		return errors.New("パスワードに数字を入れてください")
	} else if passwordLength < 8 {
		return errors.New("8文字以上のパスワードを入力してください")
	} else if passwordLength > 64 {
		return errors.New("64文字以下のパスワードを入力してください")
	}

	return nil
}
