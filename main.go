package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main.go <length>")
		return
	}
	length := os.Args[1]

	val, err := Validate(length)
	if err != nil {
		fmt.Printf("Error with value input: %v\n", err)
		return
	}

	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	password, err := GeneratePassword(charset, val)
	if err != nil {
		fmt.Printf("Error with password generation: %v\n", err)
		return
	}

	fmt.Printf("Password: %q\n", password)
}

func Validate(length string) (int, error) {
	val, err := strconv.Atoi(length)
	if err != nil {
		return 0, fmt.Errorf("conversion error: %v", err)
	}
	if val <= 0 {
		return 0, fmt.Errorf("parameter(%v) must be a positive integer", length)
	}
	return val, nil
}