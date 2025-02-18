package models

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

type Wallet struct {
	Address string `gorm:"primaryKey;size:64;index:idx_wallet_address"`
	Balance int64  `gorm:"not null"`
}

func (w *Wallet) CreateWalletAddress() {
	w.Address = generateWalletAddress()
}

func generateWalletAddress() string {
	u := uuid.New().String()
	hash := sha256.Sum256([]byte(u))
	return hex.EncodeToString(hash[:])
}
