package logger

import (
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

func (f *LogFields) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	f.kv.ForEach(func(key string, value any) bool {
		result[key] = value
		return true // 继续迭代
	})
	return result
}

// WithField .
func (f *LogFields) WithField(key string, value any) *LogFields {
	f.kv.Set(key, value)
	return f
}

func (f *LogFields) Info(args ...interface{}) {
	DefaultLogger.Log(InfoLevel, args...)
}

func (f *LogFields) Infof(template string, args ...interface{}) {
	DefaultLogger.Fields(f.ToMap()).Logf(InfoLevel, template, args...)
}

func (f *LogFields) Trace(args ...interface{}) {
	DefaultLogger.Log(TraceLevel, args...)
}

func (f *LogFields) Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(TraceLevel, template, args...)
}

func (f *LogFields) Debug(args ...interface{}) {
	DefaultLogger.Log(DebugLevel, args...)
}

func (f *LogFields) Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(DebugLevel, template, args...)
}

func (f *LogFields) Warn(args ...interface{}) {
	DefaultLogger.Log(WarnLevel, args...)
}

func (f *LogFields) Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(WarnLevel, template, args...)
}
