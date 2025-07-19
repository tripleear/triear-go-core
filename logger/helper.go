package logger

import (
	"context"
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

func (h *Helper) toFieldsMap() map[string]any {
	result := make(map[string]any)
	if h.kv != nil {
		h.kv.ForEach(func(key string, value any) bool {
			result[key] = value
			return true // 继续迭代
		})
	}
	return result
}

func (h *Helper) copyFields() *haxmap.Map[string, any] {
	r := haxmap.New[string, any]()
	if h.kv != nil {
		h.kv.ForEach(func(k string, v any) bool {
			r.Set(k, v)
			return true
		})
	}
	return r
}

func (h *Helper) Info(ctx context.Context, args ...any) {
	h.Infof(ctx, "%+v", args...)
}

func (h *Helper) Infof(ctx context.Context, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).infof(ctx, format, h.kv, args...)
}

func (h *Helper) Trace(ctx context.Context, args ...any) {
	h.Tracef(ctx, "%+v", args...)
}

func (h *Helper) Tracef(ctx context.Context, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).debugf(ctx, format, h.kv, args...)
}

func (h *Helper) Debug(ctx context.Context, args ...any) {
	h.Debugf(ctx, "%+v", args...)
}

func (h *Helper) Debugf(ctx context.Context, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).debugf(ctx, format, h.kv, args...)
}

func (h *Helper) Warn(ctx context.Context, args ...any) {
	h.Warnf(ctx, "%+v", args...)
}

func (h *Helper) Warnf(ctx context.Context, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).warnf(ctx, format, h.kv, args...)
}

func (h *Helper) Error(ctx context.Context, err error, args ...any) {
	h.Errorf(ctx, err, "%+v", args...)
}

func (h *Helper) Errorf(ctx context.Context, err error, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).errorf(ctx, err, format, h.kv, args...)
}

func (h *Helper) Fatalf(ctx context.Context, err error, format string, args ...any) {
	h.Logger.Native().(*defaultLogger).fatalf(ctx, err, format, h.kv, args...)
	os.Exit(1)
}

func (h *Helper) WithError(err error) *Helper {
	fields := h.copyFields()
	fields.Set("error", err)
	return &Helper{Logger: h.Logger, kv: fields}
}

func (h *Helper) WithFields(fields map[string]any) *Helper {
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
	return h.WithFields(map[string]any{
		"func": funcName,
	})
}
