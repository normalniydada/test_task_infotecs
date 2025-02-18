package storage

import (
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"gorm.io/gorm"
)

func GetLastNTransactions(db *gorm.DB, count int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := db.Order("created_at desc").Limit(count).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
