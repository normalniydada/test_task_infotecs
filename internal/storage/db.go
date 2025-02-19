// Package storage отвечает за подключение к базе данных и управления
package storage

import (
	"fmt"
	"github.com/normalniydada/test_task_infotecs/internal/config"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB устанавливает соединение с базой данных PostgreSQL, выполняет миграции и возвращает объект GORM
//
// Параметры:
//   - cfg (*config.DatabaseConfig): конфигурация базы данных
//   - zLog (*zap.Logger): логгер для записи информации о подключении
//
// Возвращает:
//   - *gorm.DB: объект подключения к базе данных
//
// Возможные ошибки:
//   - Завершает работу приложения (`zLog.Fatal`), если не удалось подключиться к базе данных
//   - Завершает работу приложения, если произошла ошибка при миграции таблиц
//
// Логика работы:
//  1. Формирование строки подключения DSN к базе данных.
//  2. Открытие соединения с базой данных через GORM
//  3. Выполнение автоматические миграции (`AutoMigrate`) для таблиц `Wallet` и `Transaction`
//  4. Логирование успешного подключение и миграции
func InitDB(cfg *config.DatabaseConfig, zLog *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSlMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zLog.Fatal("Database connection error: ", zap.Error(err))
	}

	zLog.Info("Database connection success",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
	)

	// Миграция
	if err = db.AutoMigrate(&models.Wallet{}, &models.Transaction{}); err != nil {
		zLog.Fatal("Database migration error: ", zap.Error(err))
	}
	zLog.Info("Database migration success")

	return db
}

// CloseDB закрывает соединение с базой данных
//
// Параметры:
//   - db (*gorm.DB): объект подключения к базе данных
//   - zLog (*zap.Logger): логгер для записи информации
//
// Возможные ошибки:
//   - Логирует ошибку, если не удалось получить объект `*sql.DB`
//   - Логирует ошибку, если не удалось корректно закрыть соединение
//
// Логика работы:
//  1. Получение `*sql.DB` из `gorm.DB`
//  2. Закрытие соединения с базой данных
//  3. Логирование успешного завершение операции
func CloseDB(db *gorm.DB, zLog *zap.Logger) {
	pdb, err := db.DB()
	if err != nil {
		zLog.Fatal("Error getting database: ", zap.Error(err))
		return
	}

	if err = pdb.Close(); err != nil {
		zLog.Fatal("Error closing database: ", zap.Error(err))
	} else {
		zLog.Info("Closing database success")
	}
}

/*func Truncate(db *gorm.DB) error {
	if err := db.Exec("TRUNCATE TABLE wallets, transactions RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}
	log.Println("[DB] Tables clear!")
	return nil
}*/
