// Package dto содержит структуры для передачи данных в API
package dto

// TransactionRequest представляет тело запроса для перевода средств
//
// Используется в API `POST /api/send`.
//
// Поля:
//   - From (string) — адрес кошелька отправителя
//   - To (string) — адрес кошелька получателя
//   - Amount (float64) — сумма перевода в у.е
//
// Пример JSON-запроса:
//
//	{
//	  "from": "wallet1",
//	  "to": "wallet2",
//	  "amount": 33.3
//	}
type TransactionRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
