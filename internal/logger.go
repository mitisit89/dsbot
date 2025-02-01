package dsbot

import (
	"log/slog"
	"os"
)

func SetUpLogger() *slog.Logger {
	var logger *slog.Logger

	if os.Getenv("LOG_LEVEL") == "debug" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
