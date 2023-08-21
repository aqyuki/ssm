package config

import (
	"os"
	"path/filepath"
)

// GetConfigureDirectory return path to save configuration
func GetConfigureDirectory() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", nil
	}
	return filepath.Join(path, "ssm"), nil
}

// ExistDir return if exist directory
func ExistDir(path string) bool {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}

// CreateStream create stream to load configuration file
func CreateStream(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
