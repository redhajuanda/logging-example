package logger

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// LoggerWrapper represent common interface for logging function
type LoggerWrapper interface {
	With(ctx context.Context, args ...interface{}) LoggerWrapper
	WithProcess(process func()) LoggerWrapper
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

func runProcess(command func()) {
	if command != nil {
		command()
	}
}
