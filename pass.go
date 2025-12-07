package main

import (
	"crypto/rand"
	mathrand "math/rand"
	"fmt"
	"math/big"
)

func GenRndRune(charset []rune) (rune, error) {
	bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, fmt.Errorf("error from rand.Int dependency %v", err)
	}
	idx := int(bigIdx.Int64())
	result := charset[idx]
	return result, nil
}

func GenPwd(options Options) (string, error) {
	charset := make(map[rune]bool)
	var charsetCounter uint
	var result []rune
	if options.includeNumbers {
		charsetNumString := "0123456789"
		for i := range charsetNumString {
			charset[rune(charsetNumString[i])] = false
		}
		n, err := GenRndRune([]rune(charsetNumString))
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)
		delete(charset, n)
		charsetCounter++
	}
	if options.includeLower {
		charsetLower := []rune("abcdefghijklmnopqrstuvwxyz")
		for i := range charsetLower {
			charset[rune(charsetLower[i])] = false
		}
		n, err := GenRndRune([]rune(charsetLower))
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)
		delete(charset, n)
		charsetCounter++
	}
	if options.includeUpper {
		charsetUpper := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		for i := range charsetUpper {
			charset[rune(charsetUpper[i])] = false
		}
		n, err := GenRndRune([]rune(charsetUpper))
		if err != nil {
			return "", fmt.Errorf("error while generating numbers %v", err)
		}
		result = append(result, n)
		delete(charset, n)
		charsetCounter++
	}

	if charsetCounter > options.length {
		return "", fmt.Errorf("length is too small to satisfy one character from each set rule")
	}
	if int(options.length - charsetCounter) > len(charset) {
		return "", fmt.Errorf("length is too big to satisfy unique characters rule (max: %v)", len(charset)+int(charsetCounter))
	}
	keys := make([]rune, len(charset))
	i := 0
	for k := range charset {
			keys[i] = k
			i++
	}

	mathrand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	rest := options.length-charsetCounter
	result = append(result, keys[:rest]...)

	mathrand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return string(result), nil
}