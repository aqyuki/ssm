package config

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestExistDir(t *testing.T) {
	_, current, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(current))
	path := filepath.Join(root, "testdata", "config")

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal : success",
			args: args{
				path: path,
			},
			want: true,
		},
		{
			name: "Normal : non exist",
			args: args{
				path: filepath.Join(path, "data"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistDir(tt.args.path); got != tt.want {
				t.Errorf("ExistDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateStream(t *testing.T) {
	_, current, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(filepath.Dir(current)), "testdata")
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "Normal : success",
			filename: "config.json",
			wantErr:  false,
		},
		{
			name:     "Normal : non exist",
			filename: "__test",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateStream(filepath.Join(root, tt.filename))
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) == !tt.wantErr {
				t.Errorf("CreateStream() got = %v", got)
				return
			}
		})
	}
}
