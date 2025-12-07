package main

import (
	"strings"
	"testing"
)

func TestGenPwd(t *testing.T) {
	type testCase struct {
		options Options
		expectingError bool
	}

	type CharsetCheck struct {
		Name string
		hasSet bool
		expectedSet bool
	}

	t.Run("error when length too long", func(t *testing.T) {
		testData := []testCase{
			{
				options: Options{length: 1, includeNumbers: true, includeLower: true, includeUpper: true}, 
				expectingError: true,
			},
			{
				options: Options{length: 1, includeNumbers: false, includeLower: false, includeUpper: true}, 
				expectingError: false,
			},
			{
				options: Options{length: 62, includeNumbers: true, includeLower: true, includeUpper: true}, 
				expectingError: false,
			},
			{
				options: Options{length: 62, includeNumbers: false, includeLower: true, includeUpper: true}, 
				expectingError: true,
			},
		}

		for i, test := range testData {
			_, err := GenPwd(test.options)
			if (err != nil) != test.expectingError {
				t.Fatalf("GenPwd(): test iteration: %v; expected error: %v; received error: %v;",
					i,
					test.expectingError, 
					err != nil, 
				)
			}
		}
	})

	t.Run("correct password length", func(t *testing.T) {
		testData := []testCase{
			{
				options: Options{length: 2, includeNumbers: false, includeLower: false, includeUpper: true},
			},
			{
				options: Options{length: 3, includeNumbers: true, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 20, includeNumbers: false, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 62, includeNumbers: true, includeLower: true, includeUpper: true},
			},
		}

		for i, test := range testData {
			password, err := GenPwd(test.options)
			if (err !=nil) {
				t.Fatalf("GenPwd(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			if (len(password) != int(test.options.length)) {
				t.Fatalf("GenPwd(): test iteration: %v; expected length: %v; resulted length: %v",
					i,
					test.options.length,
					len(password),
				)
			}
		}
	})

	t.Run("all symbols are unique", func(t *testing.T) {
		testData := []testCase{
			{
				options: Options{length: 2, includeNumbers: false, includeLower: false, includeUpper: true},
			},
			{
				options: Options{length: 3, includeNumbers: true, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 20, includeNumbers: false, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 62, includeNumbers: true, includeLower: true, includeUpper: true},
			},
		}

		for i, test := range testData {
			password, err := GenPwd(test.options)
			if (err !=nil) {
				t.Fatalf("GenPwd(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			usedValues := make(map[byte]bool)
			for j:=0; j<int(test.options.length); j++ {
				if _, ok := usedValues[password[j]]; ok {
					t.Fatalf("GenPwd(): test iteration: %v; password: %v; repeated %s rune on index %v",
						i,
						password,
						string(password[j]),
						j,
					)
					break
				} else {
					usedValues[password[j]] = true
				}
			}
		}
	})

	t.Run("at least one char from each selected set", func(t *testing.T) {
		testData := []testCase{
			{
				options: Options{length: 2, includeNumbers: false, includeLower: false, includeUpper: true},
			},
			{
				options: Options{length: 3, includeNumbers: true, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 20, includeNumbers: false, includeLower: true, includeUpper: true},
			},
			{
				options: Options{length: 62, includeNumbers: true, includeLower: true, includeUpper: true},
			},
		}

		for i, test := range testData {
			password, err := GenPwd(test.options)
			if (err !=nil) {
				t.Fatalf("GenPwd(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			usedCharset := map[string]CharsetCheck{
				"numbers": {Name: "includeNumbers", expectedSet: test.options.includeNumbers },
				"lower": {Name: "includeLower", expectedSet: test.options.includeLower },
				"upper": {Name: "includeUpper", expectedSet: test.options.includeUpper },
			}
			for j:=0; j<int(test.options.length); j++ {
				if !usedCharset["numbers"].hasSet && strings.ContainsRune(NUMBERS, rune(password[j])) {
					tmp := usedCharset["numbers"]
					tmp.hasSet = true
					usedCharset["numbers"] = tmp
				} else if !usedCharset["lower"].hasSet && strings.ContainsRune(LOWER_CASE, rune(password[j])) {
					tmp := usedCharset["lower"]
					tmp.hasSet = true
					usedCharset["lower"] = tmp
				} else if !usedCharset["upper"].hasSet && strings.ContainsRune(UPPER_CASE, rune(password[j])) {
					tmp := usedCharset["upper"]
					tmp.hasSet = true
					usedCharset["upper"] = tmp
				}
			}
			for _, used := range usedCharset {
				if used.expectedSet != used.hasSet {
					t.Fatalf("GenPwd(): test iteration: %v; %v expected: %v; %v found: %v",
						i,
						used.Name,
						used.expectedSet,
						used.Name,
						used.hasSet,
					)
				}
			}
			
		}
	})
}