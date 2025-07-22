package logger

import (
	"github.com/alphadose/haxmap"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"regexp"
	"strings"
)

var simpleFormatPattern = regexp.MustCompile(`^%[\+#]*[vTtbcdoqxXUbeEfFgGsqxXp]\s*$`)

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
	// 只要 args 有内容，不为空
	if len(args) > 0 {
		return false
	}
	// 只要 fields 非 nil 且不为空
	if fields != nil && fields.Len() > 0 {
		return false
	}

	// 格式串为空，或者是单纯的格式符
	trimmed := strings.TrimSpace(format)
	return trimmed == "" || simpleFormatPattern.MatchString(trimmed)
}

func ensureStack(err error) error {
	if err == nil {
		return nil
	}
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	if _, ok := err.(stackTracer); !ok {
		return errors.WithStack(err)
	}
	return err
}
