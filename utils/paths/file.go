package paths

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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

// CopyFile 复制源文件到目标文件
func CopyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	// 确保数据刷入磁盘
	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("error syncing file: %w", err)
	}

	return nil
}

// SafeCopyFile 安全地复制源文件到目标文件，确保复制失败时不会修改目标文件
func SafeCopyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	// 创建临时文件
	tmpFile := dst + ".tmp"
	tempFile, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer tempFile.Close()

	// 复制源文件内容到临时文件
	_, err = io.Copy(tempFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying to temp file: %w", err)
	}

	// 确保所有数据写入磁盘
	err = tempFile.Sync()
	if err != nil {
		return fmt.Errorf("error syncing temp file: %w", err)
	}

	// 关闭临时文件
	err = tempFile.Close()
	if err != nil {
		return fmt.Errorf("error closing temp file: %w", err)
	}

	// 如果目标文件已经存在，先删除它
	if _, err := os.Stat(dst); err == nil {
		// 文件存在，删除目标文件
		if err = os.Remove(dst); err != nil {
			return fmt.Errorf("error removing existing destination file: %w", err)
		}
	}

	// 将临时文件重命名为目标文件
	err = os.Rename(tmpFile, dst)
	if err != nil {
		return fmt.Errorf("error renaming temp file to destination file: %w", err)
	}

	return nil
}

// GetAbsolutePath 获取文件的绝对路径
func GetAbsolutePath(relativePath string) (string, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	if exist, _ := PathExists(absPath); !exist {
		return "", fmt.Errorf("file %s does not exist", absPath)
	}
	return absPath, nil
}

// GetFileExtension 获取文件的后缀名
func GetFileExtension(filename string) string {
	return filepath.Ext(filename)
}
