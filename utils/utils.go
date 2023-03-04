package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

// IsFileExist checks if a file exists
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// CreateFile creates a file
func CreateFile(path string) error {
	// Create all directories leading up to the file
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

// ReadFile reads a file
func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text(), nil
}

// WriteFile writes content to a file
func WriteFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
