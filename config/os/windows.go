package os

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	EnvWindowsConfigDirectory = "AppData"
	EnvWindowsUserProfile     = "UserProfile"
	ApplicationName           = "ssm"
)

var (
	ErrNotExistEnvValue = errors.New("environment variable not found")
)

func GetWindowsDefaultDirectory() (string, error) {
	usr := os.Getenv(EnvWindowsUserProfile)
	if usr == "" {
		return "", ErrNotExistEnvValue
	}
	return filepath.Join(usr, ApplicationName), nil
}

func GetWindowsDefaultLogDirectory() (string, error) {
	cnf := os.Getenv(EnvWindowsConfigDirectory)
	if cnf == "" {
		return "", ErrNotExistEnvValue
	}
	return filepath.Join(cnf, ApplicationName), nil
}
