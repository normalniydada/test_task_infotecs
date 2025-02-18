// Package services содержит бизнес-логику для работы с кошельками и транзакциями
package services

import (
	"errors"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
)

// Ошибка: кошелек с указанным адресом не найден
var ErrWalletNotFound = errors.New("wallet not found")

// GetWalletBalance получает баланс кошелька по его адресу
//
// Параметры:
//   - db (*gorm.DB): подключение к базе данных
//   - address (string): адрес кошелька
//
// Возвращает:
//   - int64: баланс кошелька в минимальных единицах валюты (копейки)
//   - error: ErrWalletNotFound, если кошелек не найден; другую ошибку, если произошел сбой в БД
//
// Логика работы:
//  1. Выполняется поиск кошелька в базе данных по `address`
//  2. Если кошелек найден, возвращение его баланса
//  3. Если кошелек отсутствует, возвращение ErrWalletNotFound
//  4. В случае ошибки в базе данных - возврат ошибки
func GetWalletBalance(db *gorm.DB, address string) (int64, error) {
	var wallet models.Wallet
	if err := db.Where("address = ?", address).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrWalletNotFound
		}
		return 0, err
	}
	return wallet.Balance, nil
}
