package vcs

import (
	"fmt"
	"path"
	"sgit/pkg/conf"
	metadata "sgit/pkg/metadate"
	"sgit/pkg/repo"
	"sgit/utils/paths"
)

// full path
func AddFile(filePath string) error {
	if exist, err := paths.PathExists(filePath); !exist {
		return err
	}
	files, err := metadata.GetMetadata()
	if err != nil {
		return err
	}
	target := metadata.GetFile(filePath, &files)
	if target != nil {
		err = paths.SafeCopyFile(filePath, path.Join(conf.Conf.Repo.Path, target.GitFilename))
		if err != nil {
			return err
		}
		return repo.AddFile(conf.Conf.Repo.Path, target.GitFilename)
	}
	hashPath, err := paths.HashPath(filePath)
	if err != nil {
		return err
	}
	// gitPath := path.Join(conf.Conf.Repo.Path, hashPath)
	gitPath := hashPath
	err = paths.SafeCopyFile(filePath, path.Join(conf.Conf.Repo.Path, hashPath))
	if err != nil {
		return err
	}
	err = repo.AddFile(conf.Conf.Repo.Path, gitPath)
	if err != nil {
		return err
	}
	files = append(files, metadata.FileMetadata{
		Filename:    filePath,
		GitFilename: gitPath})
	err = metadata.SaveMetadata(files)
	if err != nil {
		return err
	}
	return nil
}

func CommitFile(filePath string, message string) error {
	files, err := metadata.GetMetadata()
	if err != nil {
		return err
	}
	target := metadata.GetFile(filePath, &files)
	if target != nil {
		return repo.CommitFile(conf.Conf.Repo.Path, message)
	} else {
		fmt.Println("File not found in metadata")
	}
	return nil
}

func DiffFile(filePath string) error {
	files, err := metadata.GetMetadata()
	if err != nil {
		return err
	}
	target := metadata.GetFile(filePath, &files)
	if target != nil {
		diff, err := repo.DiffFile(conf.Conf.Repo.Path, target.GitFilename)
		if err != nil {
			return err
		}
		fmt.Println(diff)
		return nil
	} else {
		fmt.Println("File not found in metadata")
	}
	return nil
}

func FileLog(filePath string, length int) error {
	files, err := metadata.GetMetadata()
	if err != nil {
		return err
	}
	target := metadata.GetFile(filePath, &files)
	if target != nil {
		diff, err := repo.FileLog(conf.Conf.Repo.Path, target.GitFilename, length)
		if err != nil {
			return err
		}
		fmt.Println(diff)
		return nil
	} else {
		fmt.Println("File not found in metadata")
	}
	return nil
}
