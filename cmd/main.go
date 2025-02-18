// REST API-сервер, реализованный с использованием Gin, PostgreSQL, Gorm
// Он реализует систему обработки транзакций платежной системы
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/normalniydada/test_task_infotecs/internal/config"
	"github.com/normalniydada/test_task_infotecs/internal/handlers"
	"github.com/normalniydada/test_task_infotecs/internal/seeds"
	"github.com/normalniydada/test_task_infotecs/internal/storage"
	"github.com/normalniydada/test_task_infotecs/pkg/logger"
	"go.uber.org/zap"
)

// main инициализирует и запускает HTTP-сервер
//
// Основные шаги выполнения:
//   - Инициализация логгера (`zap.Logger`)
//   - Чтение конфигурации из `config.yaml` с использованием Viper
//   - Подключение к базе данных PostgreSQL через Gorm
//   - Создание 10 тестовых кошельков (если они отсутствуют)
//   - Регистрация API-обработчиков с использованием Gin
//   - Запуск HTTP-сервера на указанном в конфигурации порту
//
// Сервер предоставляет следующие эндпоинты:
//   - POST /api/send  — отправляет средства с одного из кошельков на указанный кошелек
//   - GET  /api/transactions?count=N  — получение списка последних N транзакций
//   - GET  /api/wallet/{address}/balance  — получение баланса указанного кошелька
//
// Если сервер не может быть запущен, программа завершает выполнение с критической ошибкой
func main() {
	// Инициализация логгера
	zLog := logger.InitLogger()
	defer zLog.Sync()

	// Отключение стандартных логов Gin
	gin.SetMode(gin.ReleaseMode)

	// Загрузка конфигурации
	cfg := config.MustLoad(zLog)

	// Подключение к базе данных
	db := storage.InitDB(&cfg.Database, zLog)
	defer storage.CloseDB(db, zLog)

	// Инициализация 10 тестовых кошельков
	seeds.InitWallets(db, zLog)

	// Создание HTTP-сервера
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/send", handlers.SendTransaction(db))
		api.GET("/transactions", handlers.GetLastTransactions(db))
		api.GET("/wallet/:address/balance", handlers.GetBalance(db))
	}

	// Логирование запуска сервера
	zLog.Info("Server is running...", zap.String("address", cfg.Server.Address))

	// Запуск сервера
	if err := r.Run(cfg.Server.Address); err != nil {
		zLog.Fatal("Error start the server", zap.Error(err))
	}
}
