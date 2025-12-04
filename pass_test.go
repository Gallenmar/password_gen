package main

import "testing"

func TestGeneratePassword(t *testing.T) {
	type testCase struct {
		charset string
		length int
		expectedError bool
	}

	t.Run("error when length too long", func(t *testing.T) {
		testData := []testCase{
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 8},
			{charset: "abcdef", length: 8, expectedError: true},
		}

		for i, test := range testData {
			_, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) == !test.expectedError {
				t.Errorf("GeneratePassword(): test iteration: %v; expected: %v; received: %v;",
					i,
					test.expectedError, 
					err !=nil, 
				)
			}
		}
	})

	t.Run("correct password length", func(t *testing.T) {
		testData := []testCase{
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 8},
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 1},
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 0},
		}

		for i, test := range testData {
			password, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) {
				t.Errorf("GeneratePassword(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			if (len(password) != test.length) {
				t.Errorf("GeneratePassword(): test iteration: %v; expected length: %v; resulted length: %v",
					i,
					test.length,
					len(password),
				)
			}
		}
	})

	t.Run("all symbols are unique", func(t *testing.T) {
		testData := []testCase{
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 26},
			{charset: "abcdefghijklmnopqrstuvwxyz", length: 0},
		}

		for i, test := range testData {
			password, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) {
				t.Errorf("GeneratePassword(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			usedValues := make(map[byte]bool)
			for j:=0; j<test.length; j++ {
				if _, ok := usedValues[password[j]]; ok {
					t.Errorf("GeneratePassword(): test iteration: %v; password: %v; repeated %s rune on index %v",
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
}