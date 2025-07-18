package logger

import (
	"context"
	"github.com/rs/zerolog"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger(WithLevel(zerolog.DebugLevel), WithName("test"))
	h1 := NewHelper(l).WithFields(map[string]interface{}{"key1": "val1"})
	ctx := context.Background()
	h1.Trace(ctx, "trace_msg1")
	h1.Warn(ctx, "warn_msg1")

	h2 := NewHelper(l).WithFields(map[string]interface{}{"key2": "val2"})
	h2.Trace(ctx, "trace_msg2")
	h2.Warn(ctx, "warn_msg2")

	h3 := NewHelper(l).WithFields(map[string]interface{}{"key3": "val4"})
	h3.Info(ctx, "test_msg")
	ctx = context.WithValue(ctx, &loggerKey{}, h3)
	v := ctx.Value(&loggerKey{})
	ll := v.(*Helper)
	ll.Info(ctx, "test_msg")
}
