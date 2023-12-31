package logs

import (
	"shield/common/trace"

	"github.com/sirupsen/logrus"
)

const (
	LoggerKeyTraceID = "trace_id"
	LoggerKeySpanID  = "span_id"
	LoggerKeyLogId   = "log_id"
)

type logrusHook struct {
}

func newLogrusHook() *logrusHook {
	return new(logrusHook)
}

// Fire implements logrus.Hook.
func (h *logrusHook) Fire(entry *logrus.Entry) error {
	if entry != nil && entry.Context != nil {
		tr, ok := trace.TraceFromContext(entry.Context)
		if ok {
			entry.Data[LoggerKeyTraceID] = tr.TraceID
			entry.Data[LoggerKeySpanID] = tr.SpanID
			entry.Data[LoggerKeyLogId] = tr.LogID
		}
	}

	return nil
}

// Levels implements logrus.Hook.
func (h *logrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
