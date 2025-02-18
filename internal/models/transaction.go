// Package models содержит описание структур базы данных для работы с транзакциями и кошельками
package models

import "time"

// Transaction представляет модель транзакции между двумя кошельками
//
// Поля:
//   - ID (uint) — уникальный идентификатор транзакции (первичный ключ)
//   - From (string) — адрес кошелька отправителя (индексирован для быстрого поиска)
//   - To (string) — адрес кошелька получателя (индексирован для быстрого поиска)
//   - Amount (int64) — сумма перевода в минимальных единицах валюты (копейки)
//   - CreatedAt (time.Time) — время создания транзакции (автоматически проставляется GORM)
type Transaction struct {
	ID        uint      `gorm:"primary_key"`                // Уникальный идентификатор транзакции
	From      string    `gorm:"index:idx_transaction_from"` // Адрес отправителя
	To        string    `gorm:"index:idx_transaction_to"`   // Адрес получателя
	Amount    int64     `gorm:"not null"`                   // Сумма перевода в минимальных единицах
	CreatedAt time.Time `gorm:"autoCreateTime"`             // Дата и время создания транзакции
}
