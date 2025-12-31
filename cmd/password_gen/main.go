package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"log"
	"time"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"password_gen/internal/pass"
)

const TIMEOUT_DEFAULT = 30 // in seconds
const PASS_HISTORY_FILE_PATH = "pass.log"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// add names to return types
func TryUniquePassword(options pass.Options, hashes []string) (string, string, error) {
	var password string
	var passHash string
	var err error
	start := time.Now()
	for {
		password, err = pass.GenPwd(options)
		if err != nil {
			return "", "", fmt.Errorf("password generation: %v", err)
		}

		passHash, err = HashPassword(password)
		if err != nil {
			return "", "", fmt.Errorf("password hashing: %v", err)
		}

		var found bool
		for _, hash := range hashes {
			trimmedHash := strings.TrimSpace(hash)
			if trimmedHash == "" {
					continue
			}
			if CheckPasswordHash(password, trimmedHash) {
				found = true
				break
			}
		}
		if !found {
			break
		} else {
			elapsed := time.Since(start)
			if elapsed > time.Second * time.Duration(options.Timeout) {
				return "", "", fmt.Errorf("timeout")
			} else {
				fmt.Println("Password was detected to have already been generated. Generating another unique password...")
			}
		}
	}
	return password, passHash, nil
}

func ExtractHashes() (*os.File, []string, error) {
	f, err := os.OpenFile(PASS_HISTORY_FILE_PATH, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
			return nil, nil, fmt.Errorf("open file: %v", err)
	}
	contentBefore, err := io.ReadAll(f)
	if err != nil {
		return nil, nil, fmt.Errorf("read file: %v", err)
	}

	hashes := strings.Split(string(contentBefore), "\n")
	return f, hashes, nil
}

func ShutFile(f *os.File) {
	if f != nil {
		if closeErr := f.Close(); closeErr != nil {
			log.Printf("Error while closing file %s: %v", PASS_HISTORY_FILE_PATH, closeErr)
		}
	}
}

func CleanHistory() error {
	err := os.Remove(PASS_HISTORY_FILE_PATH)
	if err != nil {
		return fmt.Errorf("deleting file: %v", err)
	}
	return nil
}

func InitOptions() (pass.Options) {
	lengthPtr := flag.Uint("length", 0, "Length of the output (required)")
	includeNumbersPtr := flag.Bool("numbers", false, "Include numbers (0-9)")
	includeLowerPtr := flag.Bool("lower", false, "Include lowercase letters (a-z)")
	includeUpperPtr := flag.Bool("upper", false, "Include uppercase letters (A-Z)")
	timeoutPtr := flag.Uint("timeout", 0, "Time in seconds before timeout")
	cleanupPtr := flag.Bool("cleanup", false, "Ignores all other flags and cleans the history of passwords")
	flag.Parse()

	if *cleanupPtr {
		CleanHistory()
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

	// extracts hashes of previous passwords from a file
	f, hashes, err := ExtractHashes()
	if err != nil {
		ShutFile(f)
		log.Fatalf("Error while reading hashes: %v", err)
	}

	// keeps trying to generate a unique password
	password, passHash, err := TryUniquePassword(options, hashes)
	if err != nil {
		ShutFile(f)
		log.Fatalf("Error: %v", err)
	}

	// saves hash to a file to keep passwords secure and unique
	_, err = f.WriteString(passHash + "\n")
	ShutFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Password: %q\n", password)
}