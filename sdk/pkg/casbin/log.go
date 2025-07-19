package mycasbin

import (
	"context"
	"github.com/rs/zerolog"
	"log/slog"
	"sync/atomic"

	"github.com/tripleear/triear-go-core/logger"
)

// Logger is the implementation for a Logger using golang log.
type Logger struct {
	enable int32
}

func (l *Logger) LogError(err error, msg ...string) {
	slog.Error(err.Error(), msg, err)
}

// EnableLog controls whether print the message.
func (l *Logger) EnableLog(enable bool) {
	i := 0
	if enable {
		i = 1
	}
	atomic.StoreInt32(&(l.enable), int32(i))
}

// IsEnabled returns if logger is enabled.
func (l *Logger) IsEnabled() bool {
	return atomic.LoadInt32(&(l.enable)) != 0
}

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {
	var str string
	for i := range model {
		for j := range model[i] {
			str += " " + model[i][j]
		}
		str += "\n"
	}
	logger.LoggerWrapper.GetLogger().Log(context.Background(), zerolog.InfoLevel, str)
}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	logger.LoggerWrapper.GetLogger().Fields(map[string]interface{}{
		"matcher":  matcher,
		"request":  request,
		"result":   result,
		"explains": explains,
	}).Log(context.Background(), zerolog.InfoLevel, nil)
}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {
	logger.LoggerWrapper.GetLogger().Fields(map[string]interface{}{
		"roles": roles,
	})
}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {
	data := make(map[string]interface{}, len(policy))
	for k := range policy {
		data[k] = policy[k]
	}
	logger.LoggerWrapper.GetLogger().Fields(data).Log(context.Background(), zerolog.InfoLevel, nil)
}

//func (l *Logger) Print(v ...interface{}) {
//	if l.IsEnabled() {
//		logger.LoggerWrapper.GetLogger().Log(logger.InfoLevel, v...)
//	}
//}
//
//func (l *Logger) Printf(format string, v ...interface{}) {
//	if l.IsEnabled() {
//		logger.LoggerWrapper.GetLogger().Logf(logger.InfoLevel, format, v...)
//	}
//}
