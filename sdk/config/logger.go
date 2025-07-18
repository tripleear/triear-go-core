package config

import "github.com/tripleear/triear-go-core/sdk/pkg/logger"

type Logger struct {
	Type       string
	Path       string
	Level      string
	Stdout     string
	EnabledDB  bool
	Cap        uint
	DaysToKeep uint
	SentryDSN  string
}

// Setup 设置logger
func (e Logger) Setup() {
	logger.SetupLogger(
		logger.WithType(e.Type),
		logger.WithPath(e.Path),
		logger.WithLevel(e.Level),
		logger.WithStdout(e.Stdout),
		logger.WithCap(e.Cap),
		logger.WithDaysToKeep(e.DaysToKeep),
		logger.WithSentryDSN(e.SentryDSN),
	)
}

var LoggerConfig = new(Logger)
