package logging

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Panic(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)

	With(fields ...zapcore.Field) *zap.Logger
	Named(s string) *zap.Logger
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Sync() error
	Ctx(ctx context.Context) otelzap.LoggerWithCtx
	WithOptions(opts ...zap.Option) *otelzap.Logger
	Sugar() *otelzap.SugaredLogger
	Clone(opts ...otelzap.Option) *otelzap.Logger
	DebugContext(ctx context.Context, msg string, fields ...zapcore.Field)
	InfoContext(ctx context.Context, msg string, fields ...zapcore.Field)
	WarnContext(ctx context.Context, msg string, fields ...zapcore.Field)
	ErrorContext(ctx context.Context, msg string, fields ...zapcore.Field)

	DPanicContext(ctx context.Context, msg string, fields ...zapcore.Field)
	PanicContext(ctx context.Context, msg string, fields ...zapcore.Field)
	FatalContext(ctx context.Context, msg string, fields ...zapcore.Field)
}

func NewLogger(logger *zap.Logger, serviceName string) *otelzap.Logger {
	return otelzap.New(logger.Named(serviceName))
}
