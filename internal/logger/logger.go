package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

// env const
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// Logger interface
type Interface interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	GetLogger() *slog.Logger
}

type Logger struct {
	*slog.Logger
}

var _ Interface = (*Logger)(nil)

// New method Logger constructor
func New(env string) *Logger {
	var logLevel slog.Level

	switch env {
	case envDev:
		logLevel = slog.LevelDebug
	case envProd:
		logLevel = slog.LevelInfo
	case envLocal:
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelInfo
	}
	return &Logger{
		slog.New(
			tint.NewHandler(
				os.Stdout, &tint.Options{
					Level:      logLevel,
					TimeFormat: time.Kitchen,
				})),
	}

}
func (l *Logger) GetLogger() *slog.Logger {
	return l.Logger
}
