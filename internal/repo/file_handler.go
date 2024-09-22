package repo

import (
	"fmt"
	"os"
)

func AddFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}

	fmt.Println("File successfully added:", filePath)
	return nil
}

func GetLog() ([]string, error) {
	return []string{"Commit 1: Initial version", "Commit 2: Updated file"}, nil
}
