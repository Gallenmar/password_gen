package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	length uint
	includeNumbers bool
  includeLower bool
	includeUpper bool
}

func main() {
	lengthPtr := flag.Uint("length", 0, "Length of the output (required)")
	includeNumbersPtr := flag.Bool("numbers", false, "Include numbers (0-9)")
	includeLowerPtr := flag.Bool("lower", false, "Include lowercase letters (a-z)")
	includeUpperPtr := flag.Bool("upper", false, "Include uppercase letters (A-Z)")
	flag.Parse()

	if *lengthPtr == 0 {
		fmt.Println("Error with input: --length is a required argument.")
		flag.Usage()
		os.Exit(1)
	}
	if !*includeNumbersPtr && !*includeLowerPtr && !*includeUpperPtr {
		fmt.Print("Error: at least one character set should be included.")
		fmt.Println("Include either -numbers, -lower, or -upper flag.")
		flag.Usage()
		os.Exit(1)
	}
	// check if len is less then required sets

	options := Options{
		length: *lengthPtr,
		includeNumbers: *includeNumbersPtr,
		includeLower: *includeLowerPtr,
		includeUpper: *includeUpperPtr,
	}

	password, err := GenPwd(options)
	if err != nil {
		fmt.Printf("Error with password generation: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Password: %q\n", password)
}