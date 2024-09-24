package repo

import (
	"fmt"
	"sgit/pkg/vcs"
	"sgit/utils/paths"
)

func AddFile(filePath, message string, commit bool) error {
	absPath, err := paths.GetAbsolutePath(filePath)
	if err != nil {
		return err
	}
	err = vcs.AddFile(absPath)
	if err != nil {
		return err
	}
	if commit {
		if message == "" {
			message = "Add file " + filePath
		}
		fmt.Println("Committing file...", message)
		err = vcs.CommitFile(absPath, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetLog() ([]string, error) {
	return []string{"Commit 1: Initial version", "Commit 2: Updated file"}, nil
}
