package logger

import (
	"context"
	"testing"
	"time"

	logCore "github.com/tripleear/triear-go-core/logger"
	"gorm.io/gorm/logger"
)

func TestNew(t *testing.T) {
	l := New(logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel: logger.LogLevel(
			logCore.Wrapper.GetLogger().Options().Level),
	})
	l.Info(context.TODO(), "test")
}
