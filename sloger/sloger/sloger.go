package sloger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func New() {
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
}
