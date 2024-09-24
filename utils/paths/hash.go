package paths

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"path/filepath"
)

// HashPath 计算路径的 SHA-1 哈希值，并返回较短的哈希
func HashPath(path string) (string, error) {
	// 生成 SHA-1 哈希
	hasher := sha1.New()
	_, err := hasher.Write([]byte(path))
	if err != nil {
		return "", fmt.Errorf("failed to hash path: %w", err)
	}

	// 获取哈希值并转换为16进制字符串
	hash := hex.EncodeToString(hasher.Sum(nil))

	// 取前10位作为短哈希值
	shortHash := hash[:10]
	ext := filepath.Ext(path)
	if len(ext) > 0 {
		shortHash += ext
	}
	return shortHash, nil
}
