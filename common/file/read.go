package file

import (
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
