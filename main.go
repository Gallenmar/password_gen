package main

import (
	"fmt"
)

func main() {
	length := 10
	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	password, err := GeneratePassword(charset, length)
	if err != nil {
		fmt.Printf("Error with password generation: %v\n", err)
		return
	}

	fmt.Printf("Password: %q\n", password)
}