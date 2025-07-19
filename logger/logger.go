package logger

import (
	"context"
	"github.com/rs/zerolog"
	"sync"
)

var (
	// Wrapper logger
	Wrapper = &loggerWrapper{}
)

type loggerWrapper struct {
	mu     sync.RWMutex
	logger Logger
}

func (w *loggerWrapper) SetLogger(l Logger) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	w.logger = l
}

func (w *loggerWrapper) GetLogger() Logger {
	return w.logger
}

func (w *loggerWrapper) GetExecLogger() any {
	return w.logger.GetLogger()
}

// Logger is a generic logging interface
type Logger interface {
	// Init initialises options
	Init(options ...Option) error
	// Options The Logger options
	Options() Options
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger
	// Log writes a log entry
	Log(ctx context.Context, level zerolog.Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(ctx context.Context, level zerolog.Level, format string, v ...interface{})
	// String returns the name of logger
	String() string

	Native() any

	GetLogger() any
}

func Init(opts ...Option) error {
	return Wrapper.GetLogger().Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return Wrapper.GetLogger().Fields(fields)
}

func Log(ctx context.Context, level zerolog.Level, v ...interface{}) {
	Wrapper.GetLogger().Log(ctx, level, v...)
}

func Logf(ctx context.Context, level zerolog.Level, format string, v ...interface{}) {
	Wrapper.GetLogger().Logf(ctx, level, format, v...)
}

func String() string {
	return Wrapper.GetLogger().String()
}
