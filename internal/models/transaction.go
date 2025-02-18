package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primary_key"`
	From      string    `gorm:"index:idx_transaction_from"`
	To        string    `gorm:"index:idx_transaction_to"`
	Amount    int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
