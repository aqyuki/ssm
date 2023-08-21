package logger

import (
	"io"

	"golang.org/x/exp/slog"
)

// Setup sets the output destination for slog.
func Setup(w io.Writer) {
	printer := slog.New(slog.NewJSONHandler(w, nil))
	slog.SetDefault(printer)
}
