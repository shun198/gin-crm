package config

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func TokenGenerator(length int) (string, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
