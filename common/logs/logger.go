package logs

import (
	"context"
	"io"
)

type Logger interface {
	SetLevel(level Level)
	SetOutput(io.Writer)

	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})

	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	CtxTrace(ctx context.Context, v ...interface{})
	CtxDebug(ctx context.Context, v ...interface{})
	CtxInfo(ctx context.Context, v ...interface{})
	CtxWarn(ctx context.Context, v ...interface{})
	CtxError(ctx context.Context, v ...interface{})
	CtxFatal(ctx context.Context, v ...interface{})

	CtxTracef(ctx context.Context, format string, v ...interface{})
	CtxDebugf(ctx context.Context, format string, v ...interface{})
	CtxInfof(ctx context.Context, format string, v ...interface{})
	CtxWarnf(ctx context.Context, format string, v ...interface{})
	CtxErrorf(ctx context.Context, format string, v ...interface{})
	CtxFatalf(ctx context.Context, format string, v ...interface{})
}

const defaultSkip = 3

func init() {
	defaultLogger = NewLogrusLogger()
}

var defaultLogger Logger

func GetDefaultLogger() Logger {
	if defaultLogger == nil {
		defaultLogger = NewLogrusLogger()
	}
	return defaultLogger
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
	GetDefaultLogger().SetLevel(level)
}

func Trace(v ...interface{}) {
	GetDefaultLogger().Trace(v...)
}

func Debug(v ...interface{}) {
	GetDefaultLogger().Debug(v...)
}

func Info(v ...interface{}) {
	GetDefaultLogger().Info(v...)
}

func Warn(v ...interface{}) {
	GetDefaultLogger().Warn(v...)
}

func Error(v ...interface{}) {
	GetDefaultLogger().Error(v...)
}

func Fatal(v ...interface{}) {
	GetDefaultLogger().Fatal(v...)
}

func Tracef(format string, v ...interface{}) {
	GetDefaultLogger().Tracef(format, v...)
}

func Debugf(format string, v ...interface{}) {
	GetDefaultLogger().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	GetDefaultLogger().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	GetDefaultLogger().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	GetDefaultLogger().Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	GetDefaultLogger().Fatalf(format, v...)
}

func CtxTrace(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxTrace(ctx, v...)
}

func CtxDebug(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxDebug(ctx, v...)
}

func CtxInfo(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxInfo(ctx, v...)
}

func CtxWarn(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxWarn(ctx, v...)
}

func CtxError(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxError(ctx, v...)
}

func CtxFatal(ctx context.Context, v ...interface{}) {
	GetDefaultLogger().CtxFatal(ctx, v...)
}

func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxTracef(ctx, format, v...)
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxDebugf(ctx, format, v...)
}

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxInfof(ctx, format, v...)
}

func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxWarnf(ctx, format, v...)
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxErrorf(ctx, format, v...)
}

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	GetDefaultLogger().CtxFatalf(ctx, format, v...)
}
