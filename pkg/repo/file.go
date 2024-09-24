package repo

import (
	"bytes"
	"fmt"
	"os/exec"
	"sgit/pkg/conf"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func AddFile(repoPath string, filePath string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = w.Add(filePath)
	if err != nil {
		return fmt.Errorf("failed to add file: %w", err)
	}

	fmt.Printf("Added file %s to the repository\n", filePath)
	return nil
}

func CommitFile(repoPath string, message string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	commit, err := w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  conf.Conf.Git.Name,
			Email: conf.Conf.Git.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	obj, err := r.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("failed to get commit object: %w", err)
	}

	fmt.Printf("Committed with hash: %s\n", obj.Hash)
	return nil
}

// DiffFile returns the diff of the specified file in the repository as a string.
func DiffFile(repoPath string, filePath string) (string, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return "", fmt.Errorf("failed to open repository: %w", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree: %w", err)
	}

	// Check if the file has been modified
	status, err := wt.Status()
	if err != nil {
		return "", fmt.Errorf("failed to get worktree status: %w", err)
	}

	fileStatus := status.File(filePath)
	if fileStatus == nil {
		return "", fmt.Errorf("file %s not found in worktree", filePath)
	}

	// Use `git diff` command to get the diff
	var out bytes.Buffer
	cmd := exec.Command("git", "diff", filePath)
	cmd.Dir = repoPath
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}

	// Return the diff output as a string
	return out.String(), nil
}
