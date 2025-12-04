package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


func GeneratePassword(charset string, length int) (string, error) {
	charsetLength := len(charset)
	if charsetLength < length {
		return "", fmt.Errorf("desired length  %v exceeds available characters %v", length, charsetLength)
	}

	password := make([]rune, length)

	for i := range length {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error from rand.Int dependency %v", err)
		}
		password[i] = rune(charset[idx.Int64()])
	}
	result := string(password)
	return result, nil
}