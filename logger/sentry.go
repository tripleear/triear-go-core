package logger

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc/peer"
)

const (
	TracingID = iota
)

// SentryDefer .
func SentryDefer(sentryDSN string) {
	if sentryDSN == "" {
		return
	}
	defer sentry.Flush(2 * time.Second)
	if r := recover(); r != nil {
		sentry.CaptureMessage(fmt.Sprintf("%+v: %s", r, debug.Stack()))
		panic(r)
	}
}

func genGRPCTracingInfo(ctx context.Context) (tracingInfo string) {
	if ctx == nil {
		return ""
	}

	tracing := []string{}
	if p, ok := peer.FromContext(ctx); ok {
		tracing = append(tracing, p.Addr.String())
	}

	if traceID := ctx.Value(TracingID); traceID != nil {
		if tid, ok := traceID.(string); ok {
			tracing = append(tracing, tid)
		}
	}
	tracingInfo = strings.Join(tracing, "-")
	return
}

func reportToSentry(ctx context.Context, sentryDSN string, level sentry.Level, err error, format string, args ...any) {
	if sentryDSN == "" {
		return
	}
	defer sentry.Flush(2 * time.Second)
	event, extraDetails := errors.BuildSentryReport(err)
	for k, v := range extraDetails {
		event.Extra[k] = v
	}
	event.Level = level

	if msg := fmt.Sprintf(format, args...); msg != "" {
		event.Tags["message"] = msg
	}

	if tracingInfo := genGRPCTracingInfo(ctx); tracingInfo != "" {
		event.Tags["tracing"] = tracingInfo
	}

	if res := string(*sentry.CaptureEvent(event)); res != "" {
		WithFunc("log.reportToSentry").WithField("ID", res).Info(ctx, "Report to Sentry")
	}
}
