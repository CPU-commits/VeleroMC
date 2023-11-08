package utils

import (
	"os"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", nil
	}

	return string(file), nil
}

func SaveFile(fileName string, file []byte) error {
	return os.WriteFile(fileName, file, 0644)
}
