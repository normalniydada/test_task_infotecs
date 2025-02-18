// Package models содержит описание структур базы данных для работы с кошельками и транзакциями
package models

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

// Wallet представляет модель кошелька.
//
// Поля:
//   - Address (string) — уникальный адрес кошелька (первичный ключ, индексирован)
//   - Balance (int64) — баланс кошелька в минимальных единицах валюты (копейки)
type Wallet struct {
	Address string `gorm:"primaryKey;size:64;index:idx_wallet_address"` // Уникальный адрес кошелька
	Balance int64  `gorm:"not null"`                                    // Баланс кошелька
}

// CreateWalletAddress генерирует новый уникальный адрес кошелька
// и присваивает его полю Address
func (w *Wallet) CreateWalletAddress() {
	w.Address = generateWalletAddress()
}

// generateWalletAddress создаёт уникальный идентификатор для кошелька
//
// Используется UUID v4, который затем хэшируется с помощью SHA-256
// для получения 64-символьного хешированного адреса
//
// Возвращает:
//   - string: хешированный адрес кошелька
func generateWalletAddress() string {
	u := uuid.New().String()
	hash := sha256.Sum256([]byte(u))
	return hex.EncodeToString(hash[:])
}
