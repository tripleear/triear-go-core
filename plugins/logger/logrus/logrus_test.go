package logrus

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/tripleear/triear-go-core/logger"
)

func TestName(t *testing.T) {
	l := NewLogger()

	if l.String() != "logrus" {
		t.Errorf("error: name expected 'logrus' actual: %s", l.String())
	}

	t.Logf("testing logger name: %s", l.String())
}

func TestWithFields(t *testing.T) {
	l := NewLogger(logger.WithOutput(os.Stdout)).Fields(map[string]interface{}{
		"k1": "v1",
		"k2": 123456,
	})

	logger.DefaultLogger.SetLogger(
		l)

	ctx := context.Background()
	logger.Log(ctx, zerolog.InfoLevel, "testing: Info")
	logger.Logf(ctx, zerolog.InfoLevel, "testing: %s", "Infof")
}

func TestWithError(t *testing.T) {
	l := NewLogger().Fields(map[string]interface{}{"error": errors.New("boom!")})
	logger.DefaultLogger.SetLogger(
		l)
	ctx := context.Background()
	logger.Log(ctx, zerolog.ErrorLevel, "testing: error")
}

func TestWithLogger(t *testing.T) {
	// with *logrus.Logger
	l := NewLogger(WithLogger(logrus.StandardLogger())).Fields(map[string]interface{}{
		"k1": "v1",
		"k2": 123456,
	})
	logger.DefaultLogger.SetLogger(
		l)
	ctx := context.Background()
	logger.Log(ctx, zerolog.InfoLevel, "testing: with *logrus.Logger")

	// with *logrus.Entry
	el := NewLogger(WithLogger(logrus.NewEntry(logrus.StandardLogger()))).Fields(map[string]interface{}{
		"k3": 3.456,
		"k4": true,
	})
	logger.DefaultLogger.SetLogger(
		el)
	logger.Log(ctx, zerolog.InfoLevel, "testing: with *logrus.Entry")
}

func TestJSON(t *testing.T) {
	logger.DefaultLogger.SetLogger(
		NewLogger(WithJSONFormatter(&logrus.JSONFormatter{})))
	ctx := context.Background()
	logger.Logf(ctx, zerolog.InfoLevel, "test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	logger.DefaultLogger.SetLogger(
		NewLogger())
	ctx := context.Background()
	logger.Init(logger.WithLevel(zerolog.DebugLevel))
	logger.Logf(ctx, zerolog.DebugLevel, "test show debug: %s", "debug msg")

	logger.Init(logger.WithLevel(zerolog.InfoLevel))
	logger.Logf(ctx, zerolog.InfoLevel, "test non-show debug: %s", "debug msg")
}

func TestWithReportCaller(t *testing.T) {
	logger.DefaultLogger.SetLogger(
		NewLogger(ReportCaller()))
	ctx := context.Background()
	logger.Logf(ctx, zerolog.DebugLevel, "testing: %s", "WithReportCaller")
}
