package history

import (
	"fmt"
	"os"
	"io"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const PASS_HISTORY_FILE_PATH_DEFAULT = "pass.log"

func GetFilePath() string {
	if value := os.Getenv("PASSWORD_HISTORY_FILE"); value != "" {
		return value
	}
	fmt.Println("Using default file path")
	return PASS_HISTORY_FILE_PATH_DEFAULT
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ExtractHashes() (*os.File, []string, error) {
	f, err := os.OpenFile(GetFilePath(), os.O_RDWR|os.O_CREATE, 0755)
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
			log.Printf("Error while closing file %s: %v", GetFilePath(), closeErr)
		}
	}
}

func CleanHistory() error {
	err := os.Remove(GetFilePath())
	if err != nil {
		return fmt.Errorf("deleting file: %v", err)
	}
	return nil
}