package repo

import (
	"sgit/pkg/vcs"
	"sgit/utils/paths"
)

func AddFile(filePath, message string, commit bool) error {
	absPath, err := paths.GetAbsolutePath(filePath)
	if err != nil {
		return err
	}
	vcs.AddFile(absPath)
	if commit {
		if message == "" {
			message = "Add file " + filePath
		}
		vcs.CommitFile(absPath, message)
	}
	return nil
}

func GetLog() ([]string, error) {
	return []string{"Commit 1: Initial version", "Commit 2: Updated file"}, nil
}
