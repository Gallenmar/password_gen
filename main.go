package main

import (
	"fmt"
	"os"
	"time"
)

type Options struct {
	length uint
	includeNumbers bool
  includeLower bool
	includeUpper bool
}

func main() {
	options := Options{
		length: 4,
		includeNumbers: true,
		includeLower: true,
		includeUpper: true,
	}

	start := time.Now()

	password, err := GenPwd(options)
	if err != nil {
		fmt.Printf("Error with password generation: %v\n", err)
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)


	fmt.Printf("Password: %q\n", password)
}