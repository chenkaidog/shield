package logs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	skip        int
	logger      *logrus.Logger
	currentPath string
}

func NewLogrusLogger() *logrusLogger {
	l := new(logrusLogger)
	l.logger = logrus.New()
	l.logger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	absPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	l.currentPath = absPath
	l.skip = defaultSkip

	l.SetLevel(getLogLevel())
	l.logger.SetOutput(newOutput())

	l.logger.AddHook(newLogrusHook())

	return l
}

func (l *logrusLogger) newEntry() *logrus.Entry {
	_, file, line, ok := runtime.Caller(l.skip)
	if ok {
		return l.logger.WithFields(logrus.Fields{
			"location": fmt.Sprintf("%s:%d", path.Base(file), line),
		})
	}

	return l.logger.WithFields(logrus.Fields{})
}

// SetLevel implements Logger.
func (l *logrusLogger) SetLevel(level Level) {
	switch level {
	case LevelTrace:
		l.logger.SetLevel(logrus.TraceLevel)
	case LevelDebug:
		l.logger.SetLevel(logrus.DebugLevel)
	case LevelInfo:
		l.logger.SetLevel(logrus.InfoLevel)
	case LevelWarn:
		l.logger.SetLevel(logrus.WarnLevel)
	case LevelError:
		l.logger.SetLevel(logrus.ErrorLevel)
	case LevelFatal:
		l.logger.SetLevel(logrus.FatalLevel)
	}
}

func (l *logrusLogger) SetOutput(output io.Writer) {
	l.logger.SetOutput(output)
}

// CtxDebug implements Logger.
func (l *logrusLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Debugf(format, v...)
}

// CtxError implements Logger.
func (l *logrusLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Errorf(format, v...)
}

// CtxFatal implements Logger.
func (l *logrusLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Fatalf(format, v...)
}

// CtxInfo implements Logger.
func (l *logrusLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Infof(format, v...)
}

// CtxTrace implements Logger.
func (l *logrusLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Tracef(format, v...)
}

// CtxWarn implements Logger.
func (l *logrusLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.newEntry().WithContext(ctx).Warnf(format, v...)
}

// Debugf implements Logger.
func (l *logrusLogger) Debugf(format string, v ...interface{}) {
	l.newEntry().Debugf(format, v...)
}

// Errorf implements Logger.
func (l *logrusLogger) Errorf(format string, v ...interface{}) {
	l.newEntry().Errorf(format, v...)
}

// Fatalf implements Logger.
func (l *logrusLogger) Fatalf(format string, v ...interface{}) {
	l.newEntry().Fatalf(format, v...)
}

// Infof implements Logger.
func (l *logrusLogger) Infof(format string, v ...interface{}) {
	l.newEntry().Infof(format, v...)
}

// Tracef implements Logger.
func (l *logrusLogger) Tracef(format string, v ...interface{}) {
	l.newEntry().Tracef(format, v...)
}

// Warnf implements Logger.
func (l *logrusLogger) Warnf(format string, v ...interface{}) {
	l.newEntry().Warnf(format, v...)
}

// Debug implements Logger.
func (l *logrusLogger) Debug(v ...interface{}) {
	l.newEntry().Debug(v...)
}

// Error implements Logger.
func (l *logrusLogger) Error(v ...interface{}) {
	l.newEntry().Error(v...)
}

// Fatal implements Logger.
func (l *logrusLogger) Fatal(v ...interface{}) {
	l.newEntry().Fatal(v...)
}

// Info implements Logger.
func (l *logrusLogger) Info(v ...interface{}) {
	l.newEntry().Info(v...)
}

// Trace implements Logger.
func (l *logrusLogger) Trace(v ...interface{}) {
	l.newEntry().Trace(v...)
}

// Warn implements Logger.
func (l *logrusLogger) Warn(v ...interface{}) {
	l.newEntry().Warn(v...)
}

// CtxDebug implements Logger.
func (l *logrusLogger) CtxDebug(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Debug(v...)
}

// CtxError implements Logger.
func (l *logrusLogger) CtxError(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Error(v...)
}

// CtxFatal implements Logger.
func (l *logrusLogger) CtxFatal(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Fatal(v...)
}

// CtxInfo implements Logger.
func (l *logrusLogger) CtxInfo(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Info(v...)
}

// CtxTrace implements Logger.
func (l *logrusLogger) CtxTrace(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Trace(v...)
}

// CtxWarn implements Logger.
func (l *logrusLogger) CtxWarn(ctx context.Context, v ...interface{}) {
	l.newEntry().WithContext(ctx).Warn(v...)
}
