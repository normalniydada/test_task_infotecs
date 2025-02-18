package services

import (
	"errors"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrSenderNotFound   = errors.New("sender not found")
	ErrReceiverNotFound = errors.New("receiver not found")
	ErrNotEnoughMoney   = errors.New("not enough money")
	ErrSelfTransfer     = errors.New("self transfer")
	ErrInvalidAmount    = errors.New("invalid amount")
)

func TransferMoney(db *gorm.DB, from string, to string, amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if from == to {
		return ErrSelfTransfer
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var fromWallet, toWallet models.Wallet

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("address = ?", from).
			First(&fromWallet).
			Error; err != nil {
			return ErrSenderNotFound
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("address = ?", to).
			First(&toWallet).
			Error; err != nil {
			return ErrReceiverNotFound
		}

		if err := tx.Model(&fromWallet).
			Where("balance >= ?", amount).
			Update("balance", gorm.Expr("balance - ?", amount)).
			Error; err != nil {
			return ErrNotEnoughMoney
		}
		if err := tx.Model(&toWallet).
			Update("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return err
		}

		transaction := models.Transaction{
			From:   from,
			To:     to,
			Amount: amount,
		}

		return tx.Create(&transaction).Error
	})
}
