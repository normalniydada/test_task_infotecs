package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.Logger {
	env := os.Getenv("ENV")

	var cfg zap.Config
	switch env {
	case "local":
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "dev":
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "json"
	case "prod":
		cfg = zap.NewProductionConfig()
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapLog, err := cfg.Build()
	if err != nil {
		panic("Error init logger: " + err.Error())
	}

	return zapLog
}
