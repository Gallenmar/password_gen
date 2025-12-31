package history

import (
	"fmt"
	"log"
	"time"
	"strings"

	"password_gen/internal/pass"
)

func GenUnique(options pass.Options) (string, error) {
		// extracts hashes of previous passwords from a file
	f, hashes, err := ExtractHashes()
	if err != nil {
		ShutFile(f)
		return "", fmt.Errorf("error while reading hashes: %v", err)
	}

	// keeps trying to generate a unique password
	password, passHash, err := TryUniquePassword(options, hashes)
	if err != nil {
		ShutFile(f)
		return "", fmt.Errorf("Error: %v", err)
	}

	// saves hash to a file to keep passwords secure and unique
	_, err = f.WriteString(passHash + "\n")
	ShutFile(f)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	return password, nil
}

func TryUniquePassword(options pass.Options, hashes []string) (password string, passHash string, err error) {
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
				log.Printf("Password was detected to have already been generated. Generating another unique password...")
			}
		}
	}
	return password, passHash, nil
}