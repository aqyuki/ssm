package logger_test

import (
	"bytes"
	"testing"

	"github.com/aqyuki/ssm/logger"
	"golang.org/x/exp/slog"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Normal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			logger.Setup(w)
			/*
				logger.Setup()は指定したio.Writerをslogのデフォルト出力先に設定するため、
				slogを用いた出力を行うことで正常に登録されているかを確認してる
			*/
			slog.Info("test")
			if gotW := w.String(); gotW == "" {
				t.Errorf("Error : Initialize failure")
			}
		})
	}
}
