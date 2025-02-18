// Package config отвечает за загрузку конфигурации приложения
// с использованием библиотеки Viper
package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config содержит настройки сервера и базы данных
type Config struct {
	Server   ServerConfig   // Конфигурация HTTP сервера
	Database DatabaseConfig // Конфигурация базы данных
}

// ServerConfig содержит настройки HTTP сервера.
type ServerConfig struct {
	// Address - адрес сервера (по умолчанию: "localhost:8080")
	Address string `yaml:"address" env-default:"localhost:8080"`
}

// DatabaseConfig содержит параметры подключения к базе данных
type DatabaseConfig struct {
	// Host - адрес базы данных (по умолчанию: "localhost")
	Host string `yaml:"host" env-default:"localhost"`
	// Port - порт базы данных (по умолчанию: "5432")
	Port int `yaml:"port" env-default:"5432"`
	// User - имя пользователя для подключения к базе данных (по умолчанию: "postgres")
	User string `yaml:"user" env-default:"postgres"`
	// Password - пароль пользователя для подключения к базе данных
	Password string `yaml:"password"`
	//  DBName - имя базы данных
	DBName string `yaml:"dbname"`
	// SSLMode - режим SSL (по умолчанию: "disable")
	SSlMode string `yaml:"sslmode" env-default:"disable"`
}

// MustLoad загружает конфигурацию из YAML-файла и передает ее в структуру Config
//
// # Функция принимает логгер `zap.Logger` для записи ошибок при загрузке конфигурации
//
// # Если файл конфигурации отсутствует или содержит ошибки, то программа завершается с фатальной ошибкой
//
// Параметры:
//
//	-zLog - указатель на zap.Logger, используемый для логирования.
//
// Возвращает:
//
//	-*Config: указатель на загруженную конфигурацию.
func MustLoad(zLog *zap.Logger) *Config {
	viper.SetConfigName("config")          // Имя файла-конфигурации
	viper.SetConfigType("yaml")            // Формат файла-конфигурации
	viper.AddConfigPath("internal/config") // Пусть к файлу-конфигурации

	// Чтение конфигурации
	if err := viper.ReadInConfig(); err != nil {
		zLog.Fatal("Error reading config file: ", zap.Error(err))
	}

	var cfg Config
	// Декодирование YAML в структуру Config
	if err := viper.Unmarshal(&cfg); err != nil {
		zLog.Fatal("Unable to unmarshal config file", zap.Error(err))
	}

	zLog.Info("Loaded config") // Логирование успешной загрузки конфигурации

	return &cfg
}
