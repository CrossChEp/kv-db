package test_utils

import "context"

type StubLogger struct{}

func (l StubLogger) Debug(ctx context.Context, msg string) {}
func (l StubLogger) Info(ctx context.Context, msg string)  {}
func (l StubLogger) Warn(ctx context.Context, msg string)  {}
func (l StubLogger) Error(ctx context.Context, msg string) {}
func (l StubLogger) WithFields(ctx context.Context, fields map[string]interface{}) context.Context {
	return context.Background()
}
