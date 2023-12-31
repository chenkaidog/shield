package kitex

import (
	"context"
	"io"
	"shield/common/logs"

	"github.com/cloudwego/kitex/pkg/klog"
)

func NewKitexLogger() klog.FullLogger {
	return &kitexLogger{
		logger: logs.GetDefaultLogger(),
	}
}

type kitexLogger struct {
	logger logs.Logger
}

// SetLevel implements klog.FullLogger.
func (h *kitexLogger) SetLevel(level klog.Level) {
	switch level {
	case klog.LevelTrace:
		h.logger.SetLevel(logs.LevelTrace)
	case klog.LevelDebug:
		h.logger.SetLevel(logs.LevelDebug)
	case klog.LevelInfo, klog.LevelNotice:
		h.logger.SetLevel(logs.LevelInfo)
	case klog.LevelWarn:
		h.logger.SetLevel(logs.LevelWarn)
	case klog.LevelError:
		h.logger.SetLevel(logs.LevelError)
	case klog.LevelFatal:
		h.logger.SetLevel(logs.LevelFatal)
	}
}

// SetOutput implements hlog.FullLogger.
func (h *kitexLogger) SetOutput(output io.Writer) {
	h.logger.SetOutput(output)
}

// CtxDebugf implements hlog.FullLogger.
func (h *kitexLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxDebugf(ctx, format, v...)
}

// CtxErrorf implements hlog.FullLogger.
func (h *kitexLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxErrorf(ctx, format, v...)
}

// CtxFatalf implements hlog.FullLogger.
func (h *kitexLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxFatalf(ctx, format, v...)
}

// CtxInfof implements hlog.FullLogger.
func (h *kitexLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxInfof(ctx, format, v...)
}

// CtxNoticef implements hlog.FullLogger.
func (h *kitexLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxInfof(ctx, format, v...)
}

// CtxTracef implements hlog.FullLogger.
func (h *kitexLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxTracef(ctx, format, v...)
}

// CtxWarnf implements hlog.FullLogger.
func (h *kitexLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	h.logger.CtxWarnf(ctx, format, v...)
}

// Debug implements hlog.FullLogger.
func (h *kitexLogger) Debug(v ...interface{}) {
	h.logger.Debug(v...)
}

// Debugf implements hlog.FullLogger.
func (h *kitexLogger) Debugf(format string, v ...interface{}) {
	h.logger.Debugf(format, v...)
}

// Error implements hlog.FullLogger.
func (h *kitexLogger) Error(v ...interface{}) {
	h.logger.Error(v...)
}

// Errorf implements hlog.FullLogger.
func (h *kitexLogger) Errorf(format string, v ...interface{}) {
	h.logger.Errorf(format, v...)
}

// Fatal implements hlog.FullLogger.
func (h *kitexLogger) Fatal(v ...interface{}) {
	h.logger.Fatal(v...)
}

// Fatalf implements hlog.FullLogger.
func (h *kitexLogger) Fatalf(format string, v ...interface{}) {
	h.logger.Fatalf(format, v...)
}

// Info implements hlog.FullLogger.
func (h *kitexLogger) Info(v ...interface{}) {
	h.logger.Info(v...)
}

// Infof implements hlog.FullLogger.
func (h *kitexLogger) Infof(format string, v ...interface{}) {
	h.logger.Infof(format, v...)
}

// Notice implements hlog.FullLogger.
func (h *kitexLogger) Notice(v ...interface{}) {
	h.logger.Info(v...)
}

// Noticef implements hlog.FullLogger.
func (h *kitexLogger) Noticef(format string, v ...interface{}) {
	h.logger.Infof(format, v...)
}

// Trace implements hlog.FullLogger.
func (h *kitexLogger) Trace(v ...interface{}) {
	h.logger.Trace(v...)
}

// Tracef implements hlog.FullLogger.
func (h *kitexLogger) Tracef(format string, v ...interface{}) {
	h.logger.Tracef(format, v...)
}

// Warn implements hlog.FullLogger.
func (h *kitexLogger) Warn(v ...interface{}) {
	h.logger.Warn(v...)
}

// Warnf implements hlog.FullLogger.
func (h *kitexLogger) Warnf(format string, v ...interface{}) {
	h.logger.Warnf(format, v...)
}
