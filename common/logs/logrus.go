package logs

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"shield/common/constant"
	"time"

	"github.com/sirupsen/logrus"
)

const defaultSkip = 2

func init() {
	defaultLogger = NewLogger()
}

var defaultLogger *Logger

func GetDefaultLogger() *Logger {
	return defaultLogger
}

type Logger struct {
	skip int
	*logrus.Logger
	currentPath string
}

func NewLogger() *Logger {
	l := new(Logger)
	l.Logger = logrus.New()
	l.Logger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	l.currentPath = absPath
	l.setLevel(getLogLevel())
	l.SetOutput(newOutput())

	return l
}

func (l *Logger) setLevel(level Level) {
	switch level {
	case LevelTrace:
		l.SetLevel(logrus.TraceLevel)
	case LevelDebug:
		l.SetLevel(logrus.DebugLevel)
	case LevelInfo:
		l.SetLevel(logrus.InfoLevel)
	case LevelWarn:
		l.SetLevel(logrus.WarnLevel)
	case LevelError:
		l.SetLevel(logrus.ErrorLevel)
	case LevelFatal:
		l.SetLevel(logrus.FatalLevel)
	}
}

func (l *Logger) withContext(ctx context.Context) *logrus.Entry {
	entry := l.withLine()
	trace, ok := ctx.Value(constant.Trace{}).(constant.Trace)
	if ok {
		return entry.WithFields(
			logrus.Fields{
				"trace_id": trace.TraceID,
				"span_id":  trace.SpanID,
				"log_id":   trace.LogID,
			})
	}

	return entry
}

func (l *Logger) withLine() *logrus.Entry {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		return l.WithFields(logrus.Fields{
			"location": fmt.Sprintf("%s:%d", path.Base(file), line),
		})
	}

	return l.WithFields(logrus.Fields{})
}

type Level int

const (
	LevelTrace Level = iota + 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func SetLevel(level Level) {
	GetDefaultLogger().setLevel(level)
}

func Trace(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Tracef(format, v...)
}

func Debug(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	GetDefaultLogger().withLine(2).Fatalf(format, v...)
}

func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Tracef(format, v...)
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Debugf(format, v...)
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Infof(format, v...)
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Warnf(format, v...)
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Errorf(format, v...)
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().withContext(ctx).Fatalf(format, v...)
}
