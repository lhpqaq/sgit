package paths

import (
	"os"
	"path/filepath"
	"runtime"
)

// Paths holds the directory paths for config, data, cache, and logs
type Paths struct {
	ConfigDir string
	DataDir   string
	CacheDir  string
	LogDir    string
}

// GetDirectories returns the common directories based on the operating system
func GetDirectories() (*Paths, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	var paths Paths

	switch runtime.GOOS {
	case "darwin": // macOS
		paths.ConfigDir = filepath.Join(userHome)
		paths.DataDir = filepath.Join(userHome, "Library", "Application Support")
		paths.CacheDir = filepath.Join(userHome, "Library", "Caches")
		paths.LogDir = filepath.Join(userHome, "Library", "Logs")
	case "linux":
		paths.ConfigDir = filepath.Join(userHome, ".config")
		paths.DataDir = filepath.Join(userHome, ".local", "share")
		paths.CacheDir = filepath.Join(userHome, ".cache")
		paths.LogDir = "/var/log"
	case "windows":
		paths.ConfigDir = filepath.Join(userHome, "AppData", "Roaming")
		paths.DataDir = filepath.Join(userHome, "AppData", "Local")
		paths.CacheDir = filepath.Join(os.Getenv("TEMP"))
		paths.LogDir = filepath.Join(userHome, "AppData", "Local", "Logs")
	default:
		return nil, nil
	}

	return &paths, nil
}
