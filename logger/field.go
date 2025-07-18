package logger

import (
	"context"

	"github.com/alphadose/haxmap"
)

// LogFields is a wrapper for zerolog.Entry
// we need to insert some sentry captures here
type LogFields struct {
	kv *haxmap.Map[string, any]
}

// WithFunc is short for WithField
func WithFunc(fname string) *LogFields {
	return WithField("func", fname)
}

// WithField add kv into log entry
func WithField(key string, value any) *LogFields {
	r := haxmap.New[string, any]()
	r.Set(key, value)
	return &LogFields{
		kv: r,
	}
}

// WithField .
func (f *LogFields) WithField(key string, value any) *LogFields {
	f.kv.Set(key, value)
	return f
}

// Fatalf forwards to sentry
func (f LogFields) Fatalf(ctx context.Context, err error, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).fatalf(ctx, err, format, f.kv, args...)
}

// Warnf is Warnf
func (f LogFields) Warnf(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).warnf(ctx, format, f.kv, args...)
}

// Warn is Warn
func (f LogFields) Warn(ctx context.Context, args ...any) {
	f.Warnf(ctx, "%+v", args...)
}

// Infof is Infof
func (f LogFields) Infof(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).infof(ctx, format, f.kv, args...)
}

// Info is Info
func (f LogFields) Info(ctx context.Context, args ...any) {
	f.Infof(ctx, "%+v", args...)
}

// Debugf is Debugf
func (f LogFields) Debugf(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).debugf(ctx, format, f.kv, args...)
}

// Debug is Debug
func (f LogFields) Debug(ctx context.Context, args ...any) {
	f.Debugf(ctx, "%+v", args...)
}

// Errorf forwards to sentry
func (f LogFields) Errorf(ctx context.Context, err error, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).errorf(ctx, err, format, f.kv, args...)
}

// Error forwards to sentry
func (f LogFields) Error(ctx context.Context, err error, args ...any) {
	f.Errorf(ctx, err, "%+v", args...)
}
