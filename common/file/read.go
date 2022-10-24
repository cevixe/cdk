package file

import (
	"errors"
	"log"
	"os"
)

func GetFileContent(location string) string {

	buffer, err := os.ReadFile(location)
	if err != nil {
		log.Fatalf("cannot read file content: %v", err)
	}
	return string(buffer)
}

func GetBytes(location string) *[]byte {

	buffer, err := os.ReadFile(location)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	return &buffer
}

func Exists(location string) bool {
	if _, err := os.Stat(location); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
