package middleware

import (
	"context"
	"io"
	"shield/common/logs"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitLogger() {
	hlog.SetLogger(new(hertzLogger))
}

type hertzLogger struct {
}

// CtxDebugf implements hlog.FullLogger.
func (*hertzLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	logs.CtxDebug(ctx, format, v...)
}

// CtxErrorf implements hlog.FullLogger.
func (*hertzLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// CtxFatalf implements hlog.FullLogger.
func (*hertzLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// CtxInfof implements hlog.FullLogger.
func (*hertzLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// CtxNoticef implements hlog.FullLogger.
func (*hertzLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// CtxTracef implements hlog.FullLogger.
func (*hertzLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// CtxWarnf implements hlog.FullLogger.
func (*hertzLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	panic("unimplemented")
}

// Debug implements hlog.FullLogger.
func (*hertzLogger) Debug(v ...interface{}) {
	panic("unimplemented")
}

// Debugf implements hlog.FullLogger.
func (*hertzLogger) Debugf(format string, v ...interface{}) {
	panic("unimplemented")
}

// Error implements hlog.FullLogger.
func (*hertzLogger) Error(v ...interface{}) {
	panic("unimplemented")
}

// Errorf implements hlog.FullLogger.
func (*hertzLogger) Errorf(format string, v ...interface{}) {
	panic("unimplemented")
}

// Fatal implements hlog.FullLogger.
func (*hertzLogger) Fatal(v ...interface{}) {
	panic("unimplemented")
}

// Fatalf implements hlog.FullLogger.
func (*hertzLogger) Fatalf(format string, v ...interface{}) {
	panic("unimplemented")
}

// Info implements hlog.FullLogger.
func (*hertzLogger) Info(v ...interface{}) {
	panic("unimplemented")
}

// Infof implements hlog.FullLogger.
func (*hertzLogger) Infof(format string, v ...interface{}) {
	panic("unimplemented")
}

// Notice implements hlog.FullLogger.
func (*hertzLogger) Notice(v ...interface{}) {
	panic("unimplemented")
}

// Noticef implements hlog.FullLogger.
func (*hertzLogger) Noticef(format string, v ...interface{}) {
	panic("unimplemented")
}

// SetLevel implements hlog.FullLogger.
func (*hertzLogger) SetLevel(hlog.Level) {
	panic("unimplemented")
}

// SetOutput implements hlog.FullLogger.
func (*hertzLogger) SetOutput(io.Writer) {
	panic("unimplemented")
}

// Trace implements hlog.FullLogger.
func (*hertzLogger) Trace(v ...interface{}) {
	panic("unimplemented")
}

// Tracef implements hlog.FullLogger.
func (*hertzLogger) Tracef(format string, v ...interface{}) {
	panic("unimplemented")
}

// Warn implements hlog.FullLogger.
func (*hertzLogger) Warn(v ...interface{}) {
	panic("unimplemented")
}

// Warnf implements hlog.FullLogger.
func (*hertzLogger) Warnf(format string, v ...interface{}) {
	panic("unimplemented")
}
