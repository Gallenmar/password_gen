package main

import "testing"

func TestGeneratePassword(t *testing.T) {
	type testCase struct {
		charset []byte
		length int
		expectedError bool
	}

	t.Run("error when length too long", func(t *testing.T) {
		testData := []testCase{
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 8},
			{charset: []byte("abcdef"), length: 8, expectedError: true},
		}

		for i, test := range testData {
			_, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) != test.expectedError {
				t.Fatalf("GeneratePassword(): test iteration: %v; expected: %v; received: %v;",
					i,
					test.expectedError, 
					err !=nil, 
				)
			}
		}
	})

	t.Run("correct password length", func(t *testing.T) {
		testData := []testCase{
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 8},
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 1},
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 0},
		}

		for i, test := range testData {
			password, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) {
				t.Fatalf("GeneratePassword(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			if (len(password) != test.length) {
				t.Fatalf("GeneratePassword(): test iteration: %v; expected length: %v; resulted length: %v",
					i,
					test.length,
					len(password),
				)
			}
		}
	})

	t.Run("all symbols are unique", func(t *testing.T) {
		testData := []testCase{
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 26},
			{charset: []byte("abcdefghijklmnopqrstuvwxyz"), length: 0},
		}

		for i, test := range testData {
			password, err := GeneratePassword(test.charset, test.length)
			if (err !=nil) {
				t.Fatalf("GeneratePassword(): test iteration: %v; error: %v",
					i,
					err,
				)
			}
			usedValues := make(map[byte]bool)
			for j:=0; j<test.length; j++ {
				if _, ok := usedValues[password[j]]; ok {
					t.Fatalf("GeneratePassword(): test iteration: %v; password: %v; repeated %s rune on index %v",
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