package logger

import (
	"context"
	"log/slog"
	"os"
	"time"
)

var (
	defaultLogger *slog.Logger
)

func init() {
	// Configuración basada en environment
	var handler slog.Handler

	env := os.Getenv("ENV")
	if env == "production" {
		// JSON estructurado para producción
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		})
	} else {
		// Human-readable para desarrollo
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	}

	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)
}

// LogHTTPRequest estructura logs de requests HTTP.
func LogHTTPRequest(ctx context.Context, method, url string, duration time.Duration, statusCode int, err error) {
	attrs := []slog.Attr{
		slog.String("component", "http_client"),
		slog.String("method", method),
		slog.String("url", url),
		slog.Duration("duration", duration),
		slog.Int("status_code", statusCode),
	}

	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
		defaultLogger.LogAttrs(ctx, slog.LevelError, "HTTP request failed", attrs...)
	} else {
		defaultLogger.LogAttrs(ctx, slog.LevelInfo, "HTTP request completed", attrs...)
	}
}

// LogTestStep para tests estructurados.
func LogTestStep(ctx context.Context, testName, step string, metadata map[string]interface{}) {
	attrs := []slog.Attr{
		slog.String("component", "test"),
		slog.String("test_name", testName),
		slog.String("step", step),
	}

	for k, v := range metadata {
		attrs = append(attrs, slog.Any(k, v))
	}

	defaultLogger.LogAttrs(ctx, slog.LevelInfo, "Test step executed", attrs...)
}

// LogAPIOperation para operaciones específicas de API.
func LogAPIOperation(ctx context.Context, operation string, entityID string, duration time.Duration, success bool, metadata map[string]interface{}) {
	attrs := []slog.Attr{
		slog.String("component", "api"),
		slog.String("operation", operation),
		slog.String("entity_id", entityID),
		slog.Duration("duration", duration),
		slog.Bool("success", success),
	}

	for k, v := range metadata {
		attrs = append(attrs, slog.Any(k, v))
	}

	level := slog.LevelInfo
	if !success {
		level = slog.LevelError
	}

	defaultLogger.LogAttrs(ctx, level, "API operation completed", attrs...)
}
