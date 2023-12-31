package trace

import "context"

type Trace struct {
	TraceID string
	SpanID  string
	LogID   string
}

func ContextWithTrace(ctx context.Context, t Trace) context.Context {
	return context.WithValue(ctx, Trace{}, t)
}

func TraceFromContext(ctx context.Context) (Trace, bool) {
	t, ok := ctx.Value(Trace{}).(Trace)
	return t, ok
}
