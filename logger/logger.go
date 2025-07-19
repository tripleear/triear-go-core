package logger

import (
	"context"
	"github.com/rs/zerolog"
	"sync"
)

var (
	// DefaultLogger logger
	DefaultLogger Logger
	loggerMu      sync.Mutex
)

// SetDefaultLogger 只允许初始化阶段调用，支持多次Set（两次）
// 之后不再调用 Set
func SetDefaultLogger(l Logger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	DefaultLogger = l
}

// GetDefaultLogger 不加锁，假设调用时初始化已完成
func GetDefaultLogger() Logger {
	return DefaultLogger
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
	return GetDefaultLogger().Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return GetDefaultLogger().Fields(fields)
}

func Log(ctx context.Context, level zerolog.Level, v ...interface{}) {
	GetDefaultLogger().Log(ctx, level, v...)
}

func Logf(ctx context.Context, level zerolog.Level, format string, v ...interface{}) {
	GetDefaultLogger().Logf(ctx, level, format, v...)
}

func String() string {
	return GetDefaultLogger().String()
}
