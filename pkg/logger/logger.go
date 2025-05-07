package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config holds the configuration for the logger
type Config struct {
	// Level is the logging level (debug, info, warn, error)
	Level string
	// PluginName is the name of the plugin or component
	PluginName string
	// Fields are additional fields to be added to all log entries
	Fields map[string]interface{}
}

var (
	// Log is the global logger instance
	Log  *zap.SugaredLogger
	once sync.Once
)

// Init initializes the logger with the specified configuration
func Init(cfg Config) error {
	var initErr error
	once.Do(func() {
		// Parse log level
		var zapLevel zapcore.Level
		if err := zapLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
			zapLevel = zapcore.InfoLevel
		}

		// Create encoder config for console output
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		// Create a custom encoder that adds a prefix to our logs
		encoder := zapcore.NewConsoleEncoder(encoderConfig)
		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stderr),
			zapLevel,
		)

		// Prepare fields
		fields := []zap.Field{
			zap.String("plugin", cfg.PluginName),
		}

		// Add additional fields
		for k, v := range cfg.Fields {
			fields = append(fields, zap.Any(k, v))
		}

		// Create logger with additional options
		logger := zap.New(core,
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
			zap.Development(),
			zap.Fields(fields...),
		)
		Log = logger.Sugar()
	})

	return initErr
}

// WithFields returns a new logger with the given fields
func WithFields(fields map[string]interface{}) *zap.SugaredLogger {
	if Log == nil {
		return nil
	}

	zapFields := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		zapFields = append(zapFields, k, v)
	}

	return Log.With(zapFields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	if Log != nil {
		return Log.Sync()
	}
	return nil
}
