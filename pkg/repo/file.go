package repo

import (
	"fmt"
	"sgit/pkg/conf"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// AddFile 添加一个文件到仓库的索引中 (staging area)
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

// CommitFile 提交一个文件
func CommitFile(repoPath string, message string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	// 提交更改
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

	// 获取提交对象以供输出
	obj, err := r.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("failed to get commit object: %w", err)
	}

	fmt.Printf("Committed with hash: %s\n", obj.Hash)
	return nil
}

// DiffFile 比较指定文件与上次提交的差异
func DiffFile(repoPath string, filePath string) error {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return fmt.Errorf("failed to get worktree status: %w", err)
	}

	// 检查指定文件的状态
	fileStatus := status.File(filePath)
	if fileStatus == nil {
		return fmt.Errorf("file %s not found in worktree", filePath)
	}

	fmt.Printf("File %s status: %s\n", filePath, fileStatus.Staging)
	return nil
}
