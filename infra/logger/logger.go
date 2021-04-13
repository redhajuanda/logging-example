package logger

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Params type, used to pass to `WithParams`.
type Params map[string]interface{}

// LoggerWrapper represent common interface for logging function
type LoggerWrapper interface {
	With(ctx context.Context, args ...interface{}) LoggerWrapper
	WithProcess(process func()) LoggerWrapper
	WithParam(key string, value interface{}) LoggerWrapper
	WithParams(params Params) LoggerWrapper
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

type loggerWrapper struct {
	*logrus.Entry
}

// New returns a new wrapper log
func New() LoggerWrapper {
	logStore := &loggerWrapper{logrus.NewEntry(logrus.New())}
	logStore.Logger.SetFormatter(&logrus.JSONFormatter{})
	logStore.Logger.SetLevel(logrus.TraceLevel)
	return logStore
}

// The arguments will also be added to every log message generated by the logger.
func (l *loggerWrapper) With(ctx context.Context, args ...interface{}) LoggerWrapper {

	reqId := middleware.GetReqID(ctx)
	return &loggerWrapper{l.WithField("trace_id", reqId)}

}

func (l *loggerWrapper) WithProcess(process func()) LoggerWrapper {
	runProcess(process)
	return l
}

func (l *loggerWrapper) WithParam(key string, value interface{}) LoggerWrapper {
	return &loggerWrapper{l.WithField(key, value)}
}

func (l *loggerWrapper) WithParams(params Params) LoggerWrapper {
	return &loggerWrapper{l.WithFields(logrus.Fields(params))}
}

func runProcess(command func()) {
	if command != nil {
		command()
	}
}
