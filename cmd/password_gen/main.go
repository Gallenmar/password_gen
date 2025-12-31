package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	"password_gen/internal/pass"
	"password_gen/internal/history"
)

const TIMEOUT_DEFAULT = 30 // in seconds

func InitOptions() (pass.Options) {
	lengthPtr := flag.Uint("length", 0, "Length of the output (required)")
	includeNumbersPtr := flag.Bool("numbers", false, "Include numbers (0-9)")
	includeLowerPtr := flag.Bool("lower", false, "Include lowercase letters (a-z)")
	includeUpperPtr := flag.Bool("upper", false, "Include uppercase letters (A-Z)")
	timeoutPtr := flag.Uint("timeout", 0, "Time in seconds before timeout")
	cleanupPtr := flag.Bool("cleanup", false, "Ignores all other flags and cleans the history of passwords")
	flag.Parse()

	if *cleanupPtr {
		history.CleanHistory()
		os.Exit(1)
	}
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

	options := pass.Options{
		Length: *lengthPtr,
		IncludeNumbers: *includeNumbersPtr,
		IncludeLower: *includeLowerPtr,
		IncludeUpper: *includeUpperPtr,
	}

	if *timeoutPtr != 0 {
		options.Timeout = *timeoutPtr
	} else {
		options.Timeout = TIMEOUT_DEFAULT
	}

	return options
}

func main() {
	// extracts options from flags
	options := InitOptions()

  password, err := history.GenUnique(options)
	if err != nil {
		log.Fatalf("Error with generating unique password: %v", err)
	}

	fmt.Printf("Password: %q\n", password)
}