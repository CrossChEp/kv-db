package logger

import "context"

type (
	adapter interface {
		Debug(ctx context.Context, msg string)
		Info(ctx context.Context, msg string)
		Warn(ctx context.Context, msg string)
		Error(ctx context.Context, msg string)
		WithFields(ctx context.Context, fields map[string]interface{}) context.Context
	}
)
