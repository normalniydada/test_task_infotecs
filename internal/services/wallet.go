package services

import (
	"errors"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
)

var ErrWalletNotFound = errors.New("wallet not found")

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
