package logger

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Params type, used to pass to `WithParams`.
type Params map[string]interface{}

// Logger represent common interface for logging function
type Logger interface {
	With(ctx context.Context) Logger
	WithProcess(process func()) Logger
	WithParam(key string, value interface{}) Logger
	WithParams(params Params) Logger
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

type logger struct {
	*logrus.Entry
}

// New returns a new wrapper log
func New() Logger {
	logStore := &logger{logrus.NewEntry(logrus.New())}
	logStore.Logger.SetFormatter(&logrus.JSONFormatter{})
	logStore.Logger.SetLevel(logrus.TraceLevel)
	return logStore
}

// With reads requestId from context and adds to log field
func (l *logger) With(ctx context.Context) Logger {

	reqId := middleware.GetReqID(ctx)
	return &logger{l.WithField("traceId", reqId)}

}

func (l *logger) WithProcess(process func()) Logger {
	runProcess(process)
	return l
}

func (l *logger) WithParam(key string, value interface{}) Logger {
	return &logger{l.WithField(key, value)}
}

func (l *logger) WithParams(params Params) Logger {
	return &logger{l.WithFields(logrus.Fields(params))}
}

func runProcess(command func()) {
	if command != nil {
		command()
	}
}
