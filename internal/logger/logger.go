// internal/logger/logger.go
package logger

import (
	"file/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

func InitLogger(cfg *config.LogConfig) error {
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}

	var writer zapcore.WriteSyncer
	if cfg.Output == "file" {
		writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    100, // megabytes
			MaxBackups: 3,
			MaxAge:     7, // days
		})
	} else {
		writer = zapcore.AddSync(os.Stdout)
	}

	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(encoder, writer, level)
	Logger = zap.New(core, zap.AddCaller())

	return nil
}

func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}
