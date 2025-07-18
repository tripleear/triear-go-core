package logger

import (
	"github.com/alphadose/haxmap"
	"os"
)

type Helper struct {
	Logger
	kv *haxmap.Map[string, any]
}

func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

// WithFunc is short for WithField
func WithFunc(funcName string) *Helper {
	return WithField("func", funcName)
}

// WithField add kv into log entry
func WithField(key string, value any) *Helper {
	r := haxmap.New[string, any]()
	r.Set(key, value)
	return &Helper{
		kv:     r,
		Logger: DefaultLogger,
	}
}

func (h *Helper) toFieldsMap() map[string]interface{} {
	result := make(map[string]interface{})
	h.kv.ForEach(func(key string, value any) bool {
		result[key] = value
		return true // 继续迭代
	})
	return result
}

func (h *Helper) copyFields() *haxmap.Map[string, any] {
	r := haxmap.New[string, any]()
	h.kv.ForEach(func(k string, v any) bool {
		r.Set(k, v)
		return true
	})
	return r
}

func (h *Helper) Info(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(InfoLevel, args...)
}

func (h *Helper) Infof(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(InfoLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(InfoLevel, template, args...)
}

func (h *Helper) Trace(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(TraceLevel, args...)
}

func (h *Helper) Tracef(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(TraceLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(TraceLevel, template, args...)
}

func (h *Helper) Debug(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(DebugLevel, args...)
}

func (h *Helper) Debugf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(DebugLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(DebugLevel, template, args...)
}

func (h *Helper) Warn(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(WarnLevel, args...)
}

func (h *Helper) Warnf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(WarnLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(WarnLevel, template, args...)
}

func (h *Helper) Error(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(ErrorLevel, args...)
}

func (h *Helper) Errorf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(ErrorLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(ErrorLevel, template, args...)
}

func (h *Helper) Fatal(args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Log(FatalLevel, args...)
	os.Exit(1)
}

func (h *Helper) Fatalf(template string, args ...interface{}) {
	if !h.Logger.Options().Level.Enabled(FatalLevel) {
		return
	}
	h.Logger.Fields(h.toFieldsMap()).Logf(FatalLevel, template, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := h.copyFields()
	fields.Set("error", err)
	return &Helper{Logger: h.Logger, kv: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	newFields := haxmap.New[string, any]()
	for k, v := range fields {
		newFields.Set(k, v)
	}
	if h.kv != nil {
		h.kv.ForEach(func(k string, v any) bool {
			newFields.Set(k, v)
			return true
		})
	}
	return &Helper{Logger: h.Logger, kv: h.copyFields()}
}

func (h *Helper) WithFunc(funcName string) *Helper {
	return h.WithFields(map[string]interface{}{
		"func": funcName,
	})
}
