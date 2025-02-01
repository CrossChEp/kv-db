package logger

import "context"

type Logger struct {
	a adapter
}

func New(opts ...Option) *Logger {
	var o options

	for _, opt := range opts {
		opt(&o)
	}

	if o.adapter == nil {
		o.adapter = newZapAdapter(o)
	}

	return &Logger{
		a: o.adapter,
	}
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.a.Debug(ctx, msg)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.a.Info(ctx, msg)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	l.a.Warn(ctx, msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.a.Error(ctx, msg)
}

func (l *Logger) WithFields(ctx context.Context, fields map[string]interface{}) context.Context {
	return l.a.WithFields(ctx, fields)
}
