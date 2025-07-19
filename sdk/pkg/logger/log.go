package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/tripleear/triear-go-core/plugins/logger/logrus"
	"io"
	"os"

	"github.com/tripleear/triear-go-core/debug/writer"
	"github.com/tripleear/triear-go-core/logger"
	"github.com/tripleear/triear-go-core/plugins/logger/zap"
	"github.com/tripleear/triear-go-core/sdk/pkg"

	log "github.com/tripleear/triear-go-core/logger"
)

// SetupLogger 日志 cap 单位为kb
func SetupLogger(opts ...Option) logger.Logger {
	op := setDefault()
	for _, o := range opts {
		o(&op)
	}
	ctx := context.Background()
	if !pkg.PathExist(op.path) {
		err := pkg.PathCreate(op.path)
		if err != nil {
			log.Fatalf(ctx, err, "create dir error: %s", err.Error())
		}
	}
	var err error
	var output io.Writer
	switch op.stdout {
	case "file":
		output, err = writer.NewFileWriter(
			writer.WithPath(op.path),
			writer.WithCap(op.cap<<10),
		)
		if err != nil {
			log.Fatal(ctx, err, "logger setup error: %s", err.Error())
		}
	default:
		output = os.Stdout
	}
	level, err := zerolog.ParseLevel(op.level)
	if err != nil {
		log.Fatalf(ctx, err, "get logger level error, %s", err.Error())
	}
	var defaultLogger log.Logger
	switch op.driver {
	case "zap":
		defaultLogger, err = zap.NewLogger(logger.WithLevel(level), zap.WithOutput(output), zap.WithCallerSkip(2))
		if err != nil {
			log.Fatalf(ctx, err, "new zap logger error, %s", err.Error())
		}
	case "logrus":
		defaultLogger = logrus.NewLogger(logger.WithLevel(level), logger.WithOutput(output), logrus.ReportCaller())
	default:
		defaultLogger = logger.NewLogger(logger.WithLevel(level), logger.WithOutput(output))
	}
	log.SetLogger(
	defaultLogger)
	return log.DefaultLogger.GetLogger()
}
