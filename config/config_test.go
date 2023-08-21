package config_test

import (
	"testing"

	"github.com/aqyuki/ssm/config"
)

func TestAppConfig_IsEnableMultiThreads(t *testing.T) {
	tests := []struct {
		name string
		c    *config.AppConfig
		want bool
	}{
		{
			name: "Normal : return true",
			c: &config.AppConfig{
				EnableThreads: true,
			},
			want: true,
		},
		{
			name: "Normal : return false",
			c: &config.AppConfig{
				EnableThreads: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsEnableMultiThreads(); got != tt.want {
				t.Errorf("AppConfig.IsEnableMultiThreads() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConfig_GetMaximumThreads(t *testing.T) {
	tests := []struct {
		name string
		c    *config.AppConfig
		want int
	}{
		{
			name: "Normal : enable multi threads",
			c: &config.AppConfig{
				EnableThreads:  true,
				MaximumThreads: 8,
			},
			want: 8,
		},
		{
			name: "Normal : disable multi threads",
			c: &config.AppConfig{
				EnableThreads:  false,
				MaximumThreads: 8,
			},
			want: 1,
		},
		{
			name: "Normal : unsupported threads count",
			c: &config.AppConfig{
				EnableThreads:  true,
				MaximumThreads: -1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetMaximumThreads(); got != tt.want {
				t.Errorf("AppConfig.GetMaximumThreads() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConfig_GetDefaultDirectory(t *testing.T) {
	tempDir := t.TempDir()
	tests := []struct {
		name    string
		c       *config.AppConfig
		want    string
		wantErr bool
	}{
		{
			name: "Normal : exist directory",
			c: &config.AppConfig{
				DefaultDirectory: tempDir,
			},
			want:    tempDir,
			wantErr: false,
		},
		{
			name: "Error : empty",
			c: &config.AppConfig{
				DefaultDirectory: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Error : not exist directory",
			c: &config.AppConfig{
				DefaultDirectory: "_sample",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetDefaultDirectory()
			if (err != nil) != tt.wantErr {
				t.Errorf("AppConfig.GetDefaultDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AppConfig.GetDefaultDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConfig_GetDefaultLogDirectory(t *testing.T) {
	tests := []struct {
		name string
		c    *config.AppConfig
		want string
	}{
		{
			name: "Normal",
			c: &config.AppConfig{
				DefaultLogDirectory: "log",
			},
			want: "log",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetDefaultLogDirectory(); got != tt.want {
				t.Errorf("AppConfig.GetDefaultLogDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConfig_GetBaseLogFileName(t *testing.T) {
	tests := []struct {
		name string
		c    *config.AppConfig
		want string
	}{
		{
			name: "Normal : success",
			c: &config.AppConfig{
				LogFileName: "log.txt",
			},
			want: "log.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetBaseLogFileName(); got != tt.want {
				t.Errorf("AppConfig.GetBaseLogFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
