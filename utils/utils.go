package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(len int) (string, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes)[:len], nil
}
