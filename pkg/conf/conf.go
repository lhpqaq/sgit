package conf

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"sgit/utils/paths"

	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool   `mapstructure:"debug"`
	MetaDataPath string `mapstructure:"metapath"`
	Repo         RepoConfig
}

type RepoConfig struct {
	Path   string `mapstructure:"path"`
	Name   string `mapstructure:"name"`
	Remote string `mapstructure:"remote"`
}

var (
	Conf *Config
	once sync.Once
)

func Init() Config {
	dirPaths, err := paths.GetDirectories()
	if err != nil {
		log.Fatalf("Error getting directories, %s", err)
	}
	viper.SetConfigName(".sgit")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dirPaths.ConfigDir)

	if _, err := os.Stat(dirPaths.ConfigDir); os.IsNotExist(err) {
		err = os.MkdirAll(dirPaths.ConfigDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating config directory, %s", err)
		}
	}

	viper.SetDefault("debug", false)
	viper.SetDefault("metapath", filepath.Join(dirPaths.DataDir, "sgit", "metadata.json"))
	viper.SetDefault("repo.path", filepath.Join(dirPaths.DataDir, "sgit"))
	viper.SetDefault("repo.name", "repo")
	viper.SetDefault("repo.remote", "")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No config file found, using default values")
			viper.WriteConfigAs(filepath.Join(dirPaths.ConfigDir, ".sgit"))
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}

func init() {
	once.Do(func() {
		config := Init()
		Conf = &config
	})
}
