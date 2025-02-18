// Package services содержит бизнес-логику для работы с транзакциями и кошельками
package services

import (
	"errors"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Определение возможных ошибок
var (
	ErrSenderNotFound   = errors.New("sender not found")   // Ошибка: кошелек отправителя не найден
	ErrReceiverNotFound = errors.New("receiver not found") // Ошибка: кошелек получателя не найден
	ErrNotEnoughMoney   = errors.New("not enough money")   // Ошибка: недостаточно средств на балансе отправителя
	ErrSelfTransfer     = errors.New("self transfer")      // Ошибка: невозможно отправить средства самому себе
	ErrInvalidAmount    = errors.New("invalid amount")     // Ошибка: сумма перевода должна быть больше 0
)

// TransferMoney выполняет перевод средств между двумя кошельками
//
// # Функция использует GORM-транзакцию и блокировку `FOR UPDATE` для предотвращения race condition
//
// Параметры:
//   - db (*gorm.DB): подключение к базе данных
//   - from (string): адрес кошелька отправителя
//   - to (string): адрес кошелька получателя
//   - amount (int64): сумма перевода в минимальных единицах валюты (копейки)
//
// Логика работы:
//  1. Проверка, что сумма > 0 и кошельки отправителя и получателя разные
//  2. Использование `db.Transaction()`, чтобы выполнить перевод атомарно
//  3. Блокировка записи `FOR UPDATE`, чтобы избежать race condition
//  4. Проверка наличие средств у отправителя перед уменьшением баланса
//  5. Обновление балансов отправителя и получателя
//  6. Создание записи транзакции в базе данных
//  7. В случае ошибки откат изменений
func TransferMoney(db *gorm.DB, from string, to string, amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if from == to {
		return ErrSelfTransfer
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var fromWallet, toWallet models.Wallet

		// Блокировка кошелька отправителя
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("address = ?", from).
			First(&fromWallet).
			Error; err != nil {
			return ErrSenderNotFound
		}

		// Блокировка кошелька получателя
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("address = ?", to).
			First(&toWallet).
			Error; err != nil {
			return ErrReceiverNotFound
		}

		// Проверка баланс отправителя перед списанием
		if err := tx.Model(&fromWallet).
			Where("balance >= ?", amount).
			Update("balance", gorm.Expr("balance - ?", amount)).
			Error; err != nil {
			return ErrNotEnoughMoney
		}

		// Начисление средств получателю
		if err := tx.Model(&toWallet).
			Update("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return err
		}

		// Создание записи транзакции
		transaction := models.Transaction{
			From:   from,
			To:     to,
			Amount: amount,
		}

		return tx.Create(&transaction).Error
	})
}
