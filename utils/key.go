package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateKey(length int) (string, error) {
	// Create a byte slice to hold the random bytes
	bytes := make([]byte, length)

	// Generate random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode the bytes to a base64 string and trim it to the desired length
	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}
