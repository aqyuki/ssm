package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrNonInformation    = errors.New("information required for execution could not be obtained")
	ErrInitializeFailure = errors.New("configure file could not be loaded")
)

func createDefaultConfig() *AppConfig {
	config := AppConfig{
		EnableThreads:       true,
		MaximumThreads:      4,
		DefaultDirectory:    "",
		DefaultLogDirectory: "",
		LogFileName:         "log.txt",
	}
	return &config
}

func createApplicationDirectory(config *AppConfig, configDir string, configFile string) error {
	if err := os.MkdirAll(configDir, 0777); err != nil {
		return ErrInitializeFailure
	}

	f, err := os.Create(configFile)
	if err != nil {
		return ErrInitializeFailure
	}

	b, err := json.Marshal(config)
	if err != nil {
		return ErrInitializeFailure
	}

	_, err = f.Write(b)
	if err != nil {
		return ErrInitializeFailure
	}
	return nil
}

// Initialize load configuration file and parse the data to a Go structure
func Initialize() (*AppConfig, error) {

	var configDir string
	configDir, err := GetConfigureDirectory()
	if err != nil {
		return nil, ErrNonInformation
	}

	configFile := filepath.Join(configDir, "config.json")
	if ExistDir(configDir) && ExistDir(configFile) {
		if r, err := CreateStream(configFile); err != nil {
			return nil, ErrInitializeFailure
		} else {
			config, err := New(r)
			if err != nil {
				return nil, ErrInitializeFailure
			}
			return config, nil
		}
	}

	config := createDefaultConfig()
	err = createApplicationDirectory(config, configDir, configFile)
	if err != nil {
		return nil, ErrInitializeFailure
	}

	return config, nil
}
