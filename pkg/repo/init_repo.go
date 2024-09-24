package repo

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
)

// PlainInit 初始化仓库
func PlainInit(path string) error {
	_, err := git.PlainInit(path, false)
	if err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	fmt.Printf("Git repository initialized at %s\n", path)
	return nil
}
