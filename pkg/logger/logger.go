// Package logger отвечает за инициализацию логгера Zap в зависимости от окружения (local/dev/prod)
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitLogger инициализирует и возвращает логгер Zap в зависимости от переменной окружения `ENV`
//
// Поддерживаемые уровни логирования:
//   - "local" — режим разработки с логами в консоли (Debug уровень)
//   - "dev" — отладочный режим, JSON-формат логов (Debug+ уровень)
//   - "prod" — продакшен-режим, JSON-формат логов (Info+ уровень)
//   - По умолчанию используется "local"
//
// Возвращает:
//   - *zap.Logger: настроенный логгер Zap
//
// В случае ошибки инициализации вызывает `panic()`
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

	// Создание логгера
	zapLog, err := cfg.Build()
	if err != nil {
		panic("Error init logger: " + err.Error())
	}

	return zapLog
}
