package hertz

import (
	"context"
	"io"
	"shield/common/logs"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type hertzLogger struct {
	logger logs.Logger
}

func NewHertzLogger() hlog.Logger {
	return &hertzLogger{
		logger: logs.GetDefaultLogger(),
	}
}

// SetLevel implements hlog.FullLogger.
func (h *hertzLogger) SetLevel(level hlog.Level) {
	switch level {
	case hlog.LevelTrace:
		h.logger.SetLevel(logs.LevelTrace)
	case hlog.LevelDebug:
		h.logger.SetLevel(logs.LevelDebug)
	case hlog.LevelInfo, hlog.LevelNotice:
		h.logger.SetLevel(logs.LevelInfo)
	case hlog.LevelWarn:
		h.logger.SetLevel(logs.LevelWarn)
	case hlog.LevelError:
		h.logger.SetLevel(logs.LevelError)
	case hlog.LevelFatal:
		h.logger.SetLevel(logs.LevelFatal)
	}
}

// SetOutput implements hlog.FullLogger.
func (h *hertzLogger) SetOutput(output io.Writer) {
	h.logger.SetOutput(output)
}

// CtxDebugf implements hlog.FullLogger.
func (h *hertzLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxDebugf(ctx, format, v...)
}

// CtxErrorf implements hlog.FullLogger.
func (h *hertzLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxErrorf(ctx, format, v...)
}

// CtxFatalf implements hlog.FullLogger.
func (h *hertzLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxFatalf(ctx, format, v...)
}

// CtxInfof implements hlog.FullLogger.
func (h *hertzLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxInfof(ctx, format, v...)
}

// CtxNoticef implements hlog.FullLogger.
func (h *hertzLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxInfof(ctx, format, v...)
}

// CtxTracef implements hlog.FullLogger.
func (h *hertzLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxTracef(ctx, format, v...)
}

// CtxWarnf implements hlog.FullLogger.
func (h *hertzLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxWarnf(ctx, format, v...)
}

// Debug implements hlog.FullLogger.
func (h *hertzLogger) Debug(v ...interface{}) {
	h.logger.Debug(v...)
}

// Debugf implements hlog.FullLogger.
func (h *hertzLogger) Debugf(format string, v ...interface{}) {
	h.logger.Debugf(format, v...)
}

// Error implements hlog.FullLogger.
func (h *hertzLogger) Error(v ...interface{}) {
	h.logger.Error(v...)
}

// Errorf implements hlog.FullLogger.
func (h *hertzLogger) Errorf(format string, v ...interface{}) {
	h.logger.Errorf(format, v...)
}

// Fatal implements hlog.FullLogger.
func (h *hertzLogger) Fatal(v ...interface{}) {
	h.logger.Fatal(v...)
}

// Fatalf implements hlog.FullLogger.
func (h *hertzLogger) Fatalf(format string, v ...interface{}) {
	h.logger.Fatalf(format, v...)
}

// Info implements hlog.FullLogger.
func (h *hertzLogger) Info(v ...interface{}) {
	h.logger.Info(v...)
}

// Infof implements hlog.FullLogger.
func (h *hertzLogger) Infof(format string, v ...interface{}) {
	h.logger.Infof(format, v...)
}

// Notice implements hlog.FullLogger.
func (h *hertzLogger) Notice(v ...interface{}) {
	h.logger.Info(v...)
}

// Noticef implements hlog.FullLogger.
func (h *hertzLogger) Noticef(format string, v ...interface{}) {
	h.logger.Infof(format, v...)
}

// Trace implements hlog.FullLogger.
func (h *hertzLogger) Trace(v ...interface{}) {
	h.logger.Trace(v...)
}

// Tracef implements hlog.FullLogger.
func (h *hertzLogger) Tracef(format string, v ...interface{}) {
	h.logger.Tracef(format, v...)
}

// Warn implements hlog.FullLogger.
func (h *hertzLogger) Warn(v ...interface{}) {
	h.logger.Warn(v...)
}

// Warnf implements hlog.FullLogger.
func (h *hertzLogger) Warnf(format string, v ...interface{}) {
	h.logger.Warnf(format, v...)
}
