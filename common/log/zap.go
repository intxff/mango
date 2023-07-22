package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
    *zap.Logger
}

func (z *zapLogger) Info(msg string, fields ...any) {
    z.Logger.Info(msg, unwrapFields(fields...)...)
}

func (z *zapLogger) Warn(msg string, fields ...any) {
    z.Logger.Warn(msg, unwrapFields(fields...)...)
}

func (z *zapLogger) Error(msg string, fields ...any) {
    z.Logger.Error(msg, unwrapFields(fields...)...)
}

func (z *zapLogger) Debug(msg string, fields ...any) {
    z.Logger.Debug(msg, unwrapFields(fields...)...)
}

func (z *zapLogger) Fatal(msg string, fields ...any) {
    z.Logger.Fatal(msg, unwrapFields(fields...)...)
}

func unwrapFields(fields ...any) []zap.Field {
    zapFields := make([]zapcore.Field, 0, len(fields))
    for _, field := range fields {
        zapFields = append(zapFields, field.(zapcore.Field))
    }
    return zapFields
}
