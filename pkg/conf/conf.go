package conf

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"sgit/utils/paths"

	"github.com/spf13/viper"
)

type Config struct {
	Debug        bool   `mapstructure:"debug"`
	MetaDataPath string `mapstructure:"metapath"`
	Repo         RepoConfig
	Git          GitConfig
}

type RepoConfig struct {
	Path   string `mapstructure:"path"`
	Name   string `mapstructure:"name"`
	Remote string `mapstructure:"remote"`
}

type GitConfig struct {
	Name  string `mapstructure:"name"`
	Email string `mapstructure:"email"`
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
	viper.SetDefault("repo.path", filepath.Join(dirPaths.DataDir, "sgit", "repo"))
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

	if config.Git.Name == "" || config.Git.Email == "" {
		gitConfig, err := readGitConfig()
		if err != nil {
			log.Fatalf("Please configure your Git user and email %v", err)
		}

		if config.Git.Name == "" {
			config.Git.Name = gitConfig.Name
		}
		if config.Git.Email == "" {
			config.Git.Email = gitConfig.Email
		}
	}
	viper.WriteConfig()
	return config
}

func CheckGitExists() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not found in PATH")
	}
	return nil
}

func readGitConfig() (GitConfig, error) {
	err := CheckGitExists()
	if err != nil {
		return GitConfig{}, err
	}
	var gitConfig GitConfig

	nameCmd := exec.Command("git", "config", "--get", "user.name")
	nameOutput, err := nameCmd.Output()
	if err != nil {
		return gitConfig, err
	}
	gitConfig.Name = strings.TrimSpace(string(nameOutput))

	emailCmd := exec.Command("git", "config", "--get", "user.email")
	emailOutput, err := emailCmd.Output()
	if err != nil {
		return gitConfig, err
	}
	gitConfig.Email = strings.TrimSpace(string(emailOutput))

	return gitConfig, nil
}

func init() {
	once.Do(func() {
		config := Init()
		Conf = &config
	})
}
