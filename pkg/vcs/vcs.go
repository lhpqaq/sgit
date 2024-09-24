package vcs

import (
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
		repo.AddFile(conf.Conf.Repo.Path, target.GitFilename)
		return nil
	}
	hashPath, err := paths.HashPath(filePath)
	if err != nil {
		return err
	}
	gitPath := path.Join(conf.Conf.Repo.Path, hashPath)
	err = paths.SafeCopyFile(filePath, gitPath)
	if err != nil {
		return err
	}
	repo.AddFile(conf.Conf.Repo.Path, gitPath)
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
		repo.CommitFile(conf.Conf.Repo.Path, message)
		return nil
	}
	return nil
}
