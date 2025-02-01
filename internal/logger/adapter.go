package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey = struct{}

type zapAdapter struct {
	log *zap.SugaredLogger
}

func newZaoAdapter(o options) *zapAdapter {
	level := levelToZapLevel[o.level]

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		MessageKey:    "message",
		CallerKey:     "caller",
		NameKey:       "logger",
		StacktraceKey: "stackTrace",
		LevelKey:      "levelKey",
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(string(level.CapitalString()[0]))
		},
		EncodeTime: zapcore.RFC3339TimeEncoder,
	}

	var encoder zapcore.Encoder
	if o.format == TextFormat {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	if o.out == nil {
		o.out = os.Stdout
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(o.out), level)

	log := zap.New(core)
	defer log.Sync()

	return &zapAdapter{
		log: log.Sugar(),
	}
}

func (z *zapAdapter) Debug(ctx context.Context, msg string) {
	fields := z.extractFields(ctx)

	z.log.With(fields...).Debug(msg)
}

func (z *zapAdapter) Info(ctx context.Context, msg string) {
	fields := z.extractFields(ctx)

	z.log.With(fields...).Info(msg)
}

func (z *zapAdapter) Warn(ctx context.Context, msg string) {
	fields := z.extractFields(ctx)

	z.log.With(fields...).Warn(msg)
}

func (z *zapAdapter) Error(ctx context.Context, msg string) {
	fields := z.extractFields(ctx)

	z.log.With(fields...).Error(msg)
}

func (z *zapAdapter) WithFields(ctx context.Context, fields map[string]interface{}) context.Context {
	exFields, ok := ctx.Value(ctxKey{}).(map[string]interface{})
	if !ok {
		exFields = make(map[string]interface{})
	}

	mergedFields := make(map[string]interface{}, len(exFields))

	for k, v := range exFields {
		mergedFields[k] = v
	}

	for k, v := range fields {
		mergedFields[k] = v
	}

	return context.WithValue(ctx, ctxKey{}, mergedFields)
}

func (z *zapAdapter) extractFields(ctx context.Context) []interface{} {
	fields, _ := ctx.Value(ctxKey{}).(map[string]interface{})

	zf := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		zf = append(zf, zap.Any(k, v))
	}

	return zf
}
