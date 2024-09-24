package paths

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func EnsureDirExists(dirPath string) error {
	exists, err := PathExists(dirPath)
	if err != nil {
		return fmt.Errorf("error checking if directory exists: %w", err)
	}
	if !exists {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}
	return nil
}

func IsDirEmpty(dirPath string) (bool, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return false, fmt.Errorf("error reading directory: %w", err)
	}
	return len(files) == 0, nil
}
