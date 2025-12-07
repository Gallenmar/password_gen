package main

import (
	"crypto/rand"
	mathRand "math/rand"
	"fmt"
	"math/big"
)

const NUMBERS = "0123456789"
const LOWER_CASE = "abcdefghijklmnopqrstuvwxyz"
const UPPER_CASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type CharsetOption struct {
	Name string
	Chars []rune
	Include bool
}

func GenRndRune(charset []rune) (rune, int, error) {
	// returns a random rune and index of that rune in the charset
	bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, 0, fmt.Errorf("from rand.Int dependency: %v", err)
	}
	idx := int(bigIdx.Int64())
	result := charset[idx]
	return result, idx, nil
}

func GetRuneAndSet(charset []rune) ([]rune, rune, error) {
	// returns random rune and slice without that rune
	result, idx, err:= GenRndRune(charset)
	newCharset := append(charset[:idx], charset[idx + 1:]...)
	return newCharset, result, err
}

func AddCharsetMinusOne(result []rune, charset []rune, addCharset []rune) ([]rune, []rune, error) {
	// appends random rune to results and appends the rest to charset
	newCharset, n, err := GetRuneAndSet(addCharset)
	if err != nil {
		return []rune{}, []rune{}, err
	}

	result = append(result, n)
	charset = append(charset, newCharset...)
	return result, charset, nil
}

func GenPwd(options Options) (string, error) {
	var charset []rune
	var charsetCounter uint
	var result []rune
	var err error

	// build full charset based on options and save one char from each set to result
	CharsetOptions := []CharsetOption{
		{Name: "numbers", Chars: []rune(NUMBERS), Include: options.includeNumbers},
		{Name: "lower", Chars: []rune(LOWER_CASE), Include: options.includeLower},
		{Name: "upper", Chars: []rune(UPPER_CASE), Include: options.includeUpper},
	}
	for _, opt := range CharsetOptions {
		if opt.Include {
			result, charset, err = AddCharsetMinusOne(result, charset, opt.Chars)
			if err != nil {
				return "", fmt.Errorf("while adding %v: %v", opt.Name, err)
			}
			charsetCounter++
		}
	}

	// sanity check
	if charsetCounter > options.length {
		return "", fmt.Errorf("length is too small to satisfy one character from each set rule (min: %v)", charsetCounter)
	}
	if int(options.length - charsetCounter) > len(charset) {
		return "", fmt.Errorf("length is too big to satisfy unique characters rule (max: %v)", len(charset)+int(charsetCounter))
	}

	// build the rest of the password, rune by rune depleting available charset
	var char rune
	for range options.length - charsetCounter {
		charset, char, err = GetRuneAndSet(charset)
		if err != nil {
			return "", fmt.Errorf("while generating rest: %v", err)
		}
		result = append(result, char)
	}

	// shuffle first examples of each set with the rest of the password
	mathRand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return string(result), nil
}