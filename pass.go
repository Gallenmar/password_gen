package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"slices"
)

func GeneratePassword(charset []byte, length int) (string, error) {
	charsetInitialLength := len(charset)
	if charsetInitialLength < length {
		return "", fmt.Errorf("desired length  %v exceeds available characters %v", length, charsetInitialLength)
	}

	password := make([]rune, length)

	for i := range password {
		bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("error from rand.Int dependency %v", err)
		}
		idx := int(bigIdx.Int64())
		password[i] = rune(charset[idx])
		charset = slices.Delete(charset, idx, idx + 1)
	}
	result := string(password)
	return result, nil
}