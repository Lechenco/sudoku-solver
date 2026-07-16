package logging

import (
	"log/slog"
	"os"
	"strings"
)

var (
	level *slog.Level
)

func SetLevel(l slog.Level) {
	level = &l
}

func getLevel() slog.Level {
	if level != nil {
		return *level
	}

	levelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch levelStr {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	default:
		return slog.LevelError
	}
}

func LoggerFactory(name string) *slog.Logger {
	file, err := os.OpenFile("solver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)

	if err != nil {
		panic(err)
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: getLevel(),
	})
	logger := slog.New(handler)
	return logger.With("path", name)
}
