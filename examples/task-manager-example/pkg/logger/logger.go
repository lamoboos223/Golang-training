package logger

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel represents the severity level of a log message
type LogLevel string

const (
	DebugLevel LogLevel = "DEBUG"
	InfoLevel  LogLevel = "INFO"
	WarnLevel  LogLevel = "WARN"
	ErrorLevel LogLevel = "ERROR"
)

// Logger wraps zap.Logger to provide a simpler interface
type Logger struct {
	*zap.Logger
}

var (
	globalLogger *Logger
	once         sync.Once
)

// New creates a new logger instance
func New() *Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1), // Skip the wrapper function to get the actual caller
	)
	if err != nil {
		panic(err)
	}

	return &Logger{logger}
}

// WithContext adds context fields to the logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	// Extract request ID from context if available
	if requestID, ok := ctx.Value("request_id").(string); ok {
		return &Logger{l.With(zap.String("request_id", requestID))}
	}
	return l
}

// WithFields adds custom fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return &Logger{l.With(zapFields...)}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

// Info logs an info message
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

// Fatal logs a fatal message and then calls os.Exit(1)
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	return l.Logger.Sync()
}
