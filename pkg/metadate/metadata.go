package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sgit/pkg/conf"
	"sgit/utils/paths"
	"sync"
)

type FileMetadata struct {
	Filename    string `json:"filename"`
	GitFilename string `json:"git_filename"`
	Hash        string `json:"hash"`
}

var (
	once sync.Once
)

func LoadMetadata(metadataFilePath string) ([]FileMetadata, error) {
	if _, err := os.Stat(metadataFilePath); os.IsNotExist(err) {
		return []FileMetadata{}, nil
	}

	data, err := os.ReadFile(metadataFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading metadata file: %w", err)
	}
	var metadataList []FileMetadata
	err = json.Unmarshal(data, &metadataList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling metadata: %w", err)
	}
	return metadataList, nil
}

func SaveMetadataFile(metadataList []FileMetadata, metadataFilePath string) error {
	data, err := json.MarshalIndent(metadataList, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling metadata: %w", err)
	}

	tempFilePath := filepath.Join(filepath.Dir(metadataFilePath), "temp_metadata.json")

	err = os.WriteFile(tempFilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to temporary metadata file: %w", err)
	}

	err = os.Rename(tempFilePath, metadataFilePath)
	if err != nil {
		return fmt.Errorf("error replacing metadata file: %w", err)
	}

	return nil
}

func SaveMetadata(metadataList []FileMetadata) error {
	metadataFilePath := conf.Conf.MetaDataPath
	if metadataFilePath == "" {
		dirPaths, err := paths.GetDirectories()
		if err != nil {
			return fmt.Errorf("error getting directories: %w", err)
		}

		metadataFilePath = filepath.Join(dirPaths.DataDir, "metadata.json")
	}
	return SaveMetadataFile(metadataList, metadataFilePath)
}

func GetMetadata() ([]FileMetadata, error) {
	metadataFilePath := conf.Conf.MetaDataPath
	if metadataFilePath == "" {
		dirPaths, err := paths.GetDirectories()
		if err != nil {
			return nil, fmt.Errorf("error getting directories: %w", err)
		}

		metadataFilePath = filepath.Join(dirPaths.DataDir, "metadata.json")
	}
	return LoadMetadata(metadataFilePath)
}

func createFileIfNotExist(path string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	return nil
}

func GetFile(filename string, metadataList *[]FileMetadata) *FileMetadata {
	for _, metadata := range *metadataList {
		if metadata.Filename == filename {
			return &metadata
		}
	}
	return nil
}

func init() {
	once.Do(func() {
		metadataFilePath := conf.Conf.MetaDataPath
		if exist, _ := paths.PathExists(metadataFilePath); !exist {
			err := createFileIfNotExist(metadataFilePath)
			if err != nil {
				panic(fmt.Sprintf("error creating metadata file: %s",
					err))
			}
			err = SaveMetadataFile([]FileMetadata{}, metadataFilePath)
			if err != nil {
				panic(fmt.Sprintf("error saving metadata file: %s",
					err))
			}
		}
	})
}
