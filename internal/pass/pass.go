package pass

import (
	"crypto/rand"
	mathRand "math/rand"
	"fmt"
	"math/big"
)

const NUMBERS = "0123456789"
const LOWER_CASE = "abcdefghijklmnopqrstuvwxyz"
const UPPER_CASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Options struct {
	Length uint
	IncludeNumbers bool
  IncludeLower bool
	IncludeUpper bool
	Timeout uint
}

type charsetOption struct {
	Name string
	Chars []rune
	Include bool
}

func genRndRune(charset []rune) (rune, int, error) {
	// returns a random rune and index of that rune in the charset
	bigIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, 0, fmt.Errorf("from rand.Int dependency: %v", err)
	}
	idx := int(bigIdx.Int64())
	result := charset[idx]
	return result, idx, nil
}

func getRuneAndSet(charset []rune) ([]rune, rune, error) {
	// returns random rune and slice without that rune
	result, idx, err:= genRndRune(charset)
	newCharset := append(charset[:idx], charset[idx + 1:]...)
	return newCharset, result, err
}

func addCharsetMinusOne(result []rune, charset []rune, addCharset []rune) ([]rune, []rune, error) {
	// appends random rune to results and appends the rest to charset
	newCharset, n, err := getRuneAndSet(addCharset)
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
	charsetOptions := []charsetOption{
		{Name: "numbers", Chars: []rune(NUMBERS), Include: options.IncludeNumbers},
		{Name: "lower", Chars: []rune(LOWER_CASE), Include: options.IncludeLower},
		{Name: "upper", Chars: []rune(UPPER_CASE), Include: options.IncludeUpper},
	}
	for _, opt := range charsetOptions {
		if opt.Include {
			result, charset, err = addCharsetMinusOne(result, charset, opt.Chars)
			if err != nil {
				return "", fmt.Errorf("while adding %v: %v", opt.Name, err)
			}
			charsetCounter++
		}
	}

	// sanity check
	if charsetCounter > options.Length {
		return "", fmt.Errorf("length is too small to satisfy one character from each set rule (min: %v)", charsetCounter)
	}
	if int(options.Length - charsetCounter) > len(charset) {
		return "", fmt.Errorf("length is too big to satisfy unique characters rule (max: %v)", len(charset)+int(charsetCounter))
	}

	// build the rest of the password, rune by rune depleting available charset
	var char rune
	for range options.Length - charsetCounter {
		charset, char, err = getRuneAndSet(charset)
		if err != nil {
			return "", fmt.Errorf("while generating rest: %v", err)
		}
		result = append(result, char)
	}

	// shuffle first examples of each set with the rest of the password
	mathRand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return string(result), nil
}