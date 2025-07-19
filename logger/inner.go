package logger

import (
	"github.com/alphadose/haxmap"
	"github.com/rs/zerolog"
)

func argsValidate(args []any) []any {
	if len(args) > 0 {
		return args
	}
	return []any{""}
}

func wrap(f *zerolog.Event, kv *haxmap.Map[string, any]) *zerolog.Event {
	if kv == nil {
		return f
	}
	kv.ForEach(func(k string, v any) bool {
		f = f.Interface(k, v)
		return true
	})
	return f
}

func isEmptyInput(format string, fields *haxmap.Map[string, any], args ...any) bool {
	if format == "" && fields == nil && len(args) == 0 {
		return true
	}
	return false
}
