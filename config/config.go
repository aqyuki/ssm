package config

import (
	"encoding/json"
	"io"
	"os"
)

type (
	ThreadsConfig interface {
		IsEnableMultiThreads() bool // IsEnableMultiThreads return if application is allowed multi thread
		GetMaximumThreads() int     // GetMaximumThreads return the maximum number of the application how many useable threads
	}

	FileConfig interface {
		GetDefaultDirectory() (string, error) // GetDefaultDirectory return the destination folder
	}

	LogConfig interface {
		GetDefaultLogDirectory() string // GetDefaultLogDirectory return the path of default log directory
		GetBaseLogFileName() string     // GetBaseLogFileName return the name of log file
	}

	AppConfig struct {
		EnableThreads       bool   `json:"enable_threads"`     // EnableThreads is a option to use threads
		MaximumThreads      int    `json:"maximum_threads"`    // MaximumThreads is a maximum number to be used
		DefaultDirectory    string `json:"default_target_dir"` // DefaultDirectory is a path to move file
		DefaultLogDirectory string `json:"default_log_dir"`    // DefaultLogDirectory is a path to save log
		LogFileName         string `json:"log_file_name"`      // LogFileName is a name of log file
	}
)

func (c *AppConfig) IsEnableMultiThreads() bool {
	return c.EnableThreads
}

func (c *AppConfig) GetMaximumThreads() int {
	if c.MaximumThreads <= 1 || !c.EnableThreads {
		return 1
	}
	return c.MaximumThreads
}

func (c *AppConfig) GetDefaultDirectory() (string, error) {
	f, err := os.Stat(c.DefaultDirectory)
	if os.IsNotExist(err) || !f.IsDir() {
		return "", os.ErrNotExist
	}
	return c.DefaultDirectory, nil
}

func (c *AppConfig) GetDefaultLogDirectory() string {
	return c.DefaultLogDirectory
}

func (c *AppConfig) GetBaseLogFileName() string {
	return c.LogFileName
}

func New(r io.Reader) (*AppConfig, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cnf := new(AppConfig)
	err = json.Unmarshal(b, cnf)
	if err != nil {
		return nil, err
	}
	return cnf, nil
}
