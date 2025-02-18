// Package seeds содержит функции для инициализации начальных данных в базе данных
package seeds

import (
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// InitWallets создаёт 10 тестовых кошельков с балансом 100.00 у.е. (10000 в минимальных единицах валюты)
//
// # Если в базе данных уже есть кошельки, функция ничего не делает
//
// Параметры:
//   - db (*gorm.DB): подключение к базе данных GORM
//   - zLog (*zap.Logger): логгер для записи событий
//
// Процесс выполнения:
//  1. Подсчитывание количества кошельков в базе данных
//  2. Если кошельки уже существуют, завершается выполнение функции
//  3. Генерация 10 новых кошельков с уникальными адресами и балансом 10000 (100.00 у.е.)
//  4. Запись кошельков в базу данных
//  5. Логирование успешного выполнения или фатальную ошибку при записи
func InitWallets(db *gorm.DB, zLog *zap.Logger) {
	var countWallets int64
	db.Model(&models.Wallet{}).Count(&countWallets)

	// Если кошельки существуют, выход
	if countWallets > 0 {
		return
	}

	// Создание 10 кошельков
	wallets := make([]models.Wallet, 10)
	for i := 0; i < 10; i++ {
		wallet := models.Wallet{Balance: 10000} // 100 у.е
		wallet.CreateWalletAddress()            // Генерация уникального адреса
		wallets[i] = wallet
	}

	// Запись кошельков в БД
	if err := db.Create(&wallets).Error; err != nil {
		zLog.Fatal("Error init wallet: ", zap.Error(err))
	}

	zLog.Info("Init wallets successfully")
}
