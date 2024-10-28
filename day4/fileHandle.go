package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read the file: %s: %w", filename, err)
	}

	return string(data), nil
}
