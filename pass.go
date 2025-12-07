package main

import (
	"crypto/rand"
	mathrand "math/rand"
	"fmt"
	"math/big"
	"slices"
)

func GenRndRune(charset []rune) ([]rune, rune, error) {
	bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return []rune("0"), 0, fmt.Errorf("error from rand.Int dependency %v", err)
	}
	idx := int(bigIdx.Int64())
	result := charset[idx]
	newCharset := append(charset[:idx], charset[idx + 1:]...)
	return newCharset, result, nil
}

func GenPwd(options Options) (string, error) {
	var charset []rune
	var charsetCounter int
	var result []rune
	if options.includeNumbers {
		charsetNum := []rune("0123456789")
		newCharset, n, err := GenRndRune(charsetNum)
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)

		charset = slices.Concat(charset, newCharset)
		charsetCounter++
	}
	if options.includeLower {
		charsetLower := []rune("abcdefghijklmnopqrstuvwxyz")
		newCharset, n, err := GenRndRune(charsetLower)
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)

		charset = slices.Concat(charset, newCharset)
		charsetCounter++
	}
	if options.includeUpper {
		charsetUpper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		newCharset, n, err := GenRndRune(charsetUpper)
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)

		charset = slices.Concat(charset, newCharset)
		charsetCounter++
	}

	if charsetCounter > int(options.length) {
		return "", fmt.Errorf("length is too small to satisfy one character from each set rule" )
	}
	if int(options.length - uint(charsetCounter)) > len(charset) {
		return "", fmt.Errorf("length is too big to satisfy unique characters rule" )
	}
	var char rune
	var err error
	for range options.length - uint(charsetCounter) {
		charset, char, err = GenRndRune(charset)
		if err != nil {
			return "", fmt.Errorf("error while generating rest %v", err)
		}
		result = append(result, char)
	}

	mathrand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return string(result), nil
}