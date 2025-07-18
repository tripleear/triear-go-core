package logger

import (
	"context"
	"fmt"
	"github.com/alphadose/haxmap"
	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

func init() {
	lvl, err := zerolog.ParseLevel(os.Getenv("GO_ADMIN_LOG_LEVEL"))
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	DefaultLogger = NewHelper(NewLogger(WithLevel(lvl)))
}

type defaultLogger struct {
	sync.RWMutex
	opts Options
	// fields to always be logged
	fields map[string]any
	logger *zerolog.Logger
}

func (l *defaultLogger) Native() any {
	return l
}

// Init (opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Fields(fields map[string]any) Logger {
	l.Lock()
	defer l.Unlock()
	l.fields = fields
	return l
}

func (l *defaultLogger) GetFields() map[string]any {
	l.Lock()
	defer l.Unlock()
	return l.fields
}

func (l *defaultLogger) addDefaultFields(fields *haxmap.Map[string, any]) {
	if fields == nil {
		fields = haxmap.New[string, any]()
	}
	for k, v := range l.fields {
		fields.Set(k, v)
	}
}

func (l *defaultLogger) fatalf(ctx context.Context, err error, format string, fields *haxmap.Map[string, any], args ...any) {
	args = argsValidate(args)
	reportToSentry(ctx, l.opts.SentryDSN, sentry.LevelFatal, err, format, args...)
	f := l.logger.Fatal()
	l.addDefaultFields(fields)
	wrap(f, fields).Err(err).Msgf(format, args...)
}

func (l *defaultLogger) warnf(_ context.Context, format string, fields *haxmap.Map[string, any], args ...any) {
	args = argsValidate(args)
	f := l.logger.Warn()
	l.addDefaultFields(fields)
	wrap(f, fields).Msgf(format, args...)
}

func (l *defaultLogger) infof(_ context.Context, format string, fields *haxmap.Map[string, any], args ...any) {
	args = argsValidate(args)
	f := l.logger.Info()
	l.addDefaultFields(fields)
	wrap(f, fields).Msgf(format, args...)
}

func (l *defaultLogger) debugf(_ context.Context, format string, fields *haxmap.Map[string, any], args ...any) {
	args = argsValidate(args)
	f := l.logger.Debug()
	l.addDefaultFields(fields)
	wrap(f, fields).Msgf(format, args...)
}

func (l *defaultLogger) tracef(_ context.Context, format string, fields *haxmap.Map[string, any], args ...any) {
	args = argsValidate(args)
	f := l.logger.Debug()
	l.addDefaultFields(fields)
	wrap(f, fields).Msgf(format, args...)
}

func (l *defaultLogger) errorf(ctx context.Context, err error, format string, fields *haxmap.Map[string, any], args ...any) {
	if err == nil {
		return
	}
	args = argsValidate(args)
	reportToSentry(ctx, l.opts.SentryDSN, sentry.LevelError, err, format, args...)
	f := l.logger.Error()
	l.addDefaultFields(fields)
	wrap(f, fields).Stack().Err(err).Msgf(format, args...)
}

func copyFields(src map[string]any) map[string]any {
	dst := make(map[string]any, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func (l *defaultLogger) Log(ctx context.Context, level zerolog.Level, v ...any) {
	l.logf(ctx, level, "%+v", v...)
}

func (l *defaultLogger) Logf(ctx context.Context, level zerolog.Level, format string, v ...any) {
	l.logf(ctx, level, format, v...)
}

func (l *defaultLogger) logf(ctx context.Context, level zerolog.Level, format string, v ...any) {
	if int(level) == -1 {
		l.debugf(ctx, format, nil, v...)
		return
	}
	switch level {
	case zerolog.DebugLevel:
		l.debugf(ctx, format, nil, v...)
	case zerolog.InfoLevel:
		l.infof(ctx, format, nil, v...)
	case zerolog.WarnLevel:
		l.warnf(ctx, format, nil, v...)
	case zerolog.ErrorLevel:
		l.errorf(ctx, nil, format, nil, v...)
	case zerolog.FatalLevel:
		l.errorf(ctx, nil, format, nil, v...)
	case zerolog.PanicLevel:

		l.fatalf(ctx, nil, format, nil, v...)
	case zerolog.NoLevel:
		l.errorf(ctx, fmt.Errorf("no log level in coere"), format, nil, v...)
	default:
		l.errorf(ctx, fmt.Errorf("nuknown log level in coere"), format, nil, v...)
	}
}

func (l *defaultLogger) Options() Options {
	// not guard against options Context values
	l.RLock()
	opts := l.opts
	opts.Fields = copyFields(l.opts.Fields)
	l.RUnlock()
	return opts
}

// NewLogger builds a new logger based on options
func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           zerolog.InfoLevel,
		Fields:          make(map[string]any),
		Out:             os.Stderr,
		CallerSkipCount: 3,
		Context:         context.Background(),
		Name:            "",
	}

	l := &defaultLogger{opts: options}
	if err := l.Init(opts...); err != nil {
		l.Log(context.Background(), zerolog.ErrorLevel, err)
	}

	rslog := zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC822,
		}).With().Timestamp().Logger()
	rslog.Level(l.opts.Level)
	zerolog.ErrorStackMarshaler = func(err error) any {
		return errors.GetSafeDetails(err).SafeDetails
	}
	l.logger = &rslog
	// Sentry
	if l.opts.SentryDSN != "" {
		sentryDSN := l.opts.SentryDSN
		WithFunc("log.SetupLog").Infof(context.Background(), "sentry %v", sentryDSN)
		_ = sentry.Init(sentry.ClientOptions{Dsn: sentryDSN})
	}

	return l
}

func Info(ctx context.Context, args ...any) {
	Infof(ctx, "%+v", args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).infof(ctx, format, nil, args...)
}

func Trace(ctx context.Context, args ...any) {
	Tracef(ctx, "%+v", args...)
}

func Tracef(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).debugf(ctx, format, nil, args...)
}

func Debug(ctx context.Context, args ...any) {
	Debugf(ctx, "%+v", args...)
}

func Debugf(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).debugf(ctx, format, nil, args...)
}

func Warn(ctx context.Context, args ...any) {
	Warnf(ctx, "%+v", args...)
}

func Warnf(ctx context.Context, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).warnf(ctx, format, nil, args...)
}

func Error(ctx context.Context, err error, args ...any) {
	Errorf(ctx, err, "%+v", args...)
}

func Errorf(ctx context.Context, err error, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).errorf(ctx, nil, format, nil, args...)
}

func Fatal(ctx context.Context, err error, args ...any) {
	Fatalf(ctx, err, "%+v", args...)
}

func Fatalf(ctx context.Context, err error, format string, args ...any) {
	DefaultLogger.Native().(*defaultLogger).fatalf(ctx, nil, format, nil, args...)
}
