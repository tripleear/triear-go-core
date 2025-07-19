package zap

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"testing"

	"github.com/tripleear/triear-go-core/debug/writer"
	"github.com/tripleear/triear-go-core/logger"
)

func TestName(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	if l.String() != "zap" {
		t.Errorf("name is error %s", l.String())
	}

	t.Logf("test logger name: %s", l.String())
}

func TestLogf(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}

	logger.SetDefaultLogger(l)
	ctx := context.Background()
	logger.Logf(ctx, zerolog.InfoLevel, "test logf: %s", "name")
}

func TestSetLevel(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.SetDefaultLogger(l)

	ctx := context.Background()
	logger.Init(logger.WithLevel(zerolog.DebugLevel))
	l.Logf(ctx, zerolog.DebugLevel, "test show debug: %s", "debug msg")

	logger.Init(logger.WithLevel(zerolog.InfoLevel))
	l.Logf(ctx, zerolog.InfoLevel, "test non-show debug: %s", "debug msg")
}

func TestWithReportCaller(t *testing.T) {
	var err error
	defaultLogger, err := NewLogger(WithCallerSkip(0))
	logger.SetDefaultLogger(defaultLogger)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	logger.Logf(ctx, zerolog.InfoLevel, "testing: %s", "WithReportCaller")
}

func TestFields(t *testing.T) {
	l, err := NewLogger()
	if err != nil {
		t.Fatal(err)
	}
	logger.SetDefaultLogger(l.Fields(map[string]interface{}{
		"x-request-id": "123456abc",
	}))
	ctx := context.Background()
	logger.GetDefaultLogger().Log(ctx, zerolog.InfoLevel, "hello")
}

func TestFile(t *testing.T) {
	output, err := writer.NewFileWriter(writer.WithPath("testdata"), writer.WithSuffix("log"))
	if err != nil {
		t.Errorf("logger setup error: %s", err.Error())
	}
	//var err error
	ctx := context.Background()
	defaultLogger, err := NewLogger(logger.WithLevel(zerolog.DebugLevel), WithOutput(output))
	logger.SetDefaultLogger(defaultLogger)
	if err != nil {
		t.Errorf("logger setup error: %s", err.Error())
	}
	logger.SetDefaultLogger(logger.GetDefaultLogger().Fields(map[string]interface{}{
		"x-request-id": "123456abc",
	}))
	fmt.Println(logger.GetDefaultLogger())
	logger.GetDefaultLogger().Log(ctx, zerolog.DebugLevel, "hello")
}

//func TestFileKeep(t *testing.T) {
//	output, err := writer.NewFileWriter(writer.WithPath("testdata"), writer.WithSuffix("log"))
//	if err != nil {
//		t.Errorf("logger setup error: %s", err.Error())
//	}
//	//var err error
//	logger.DefaultLogger, err = NewLogger(logger.WithLevel(logger.TraceLevel), WithOutput(output))
//	if err != nil {
//		t.Errorf("logger setup error: %s", err.Error())
//	}
//
//	fmt.Println(logger.DefaultLogger)
//	logger.
//}
