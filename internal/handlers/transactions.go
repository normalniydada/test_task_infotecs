// Package handlers содержит обработчики HTTP-запросов для работы с транзакциями
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/normalniydada/test_task_infotecs/internal/models/dto"
	"github.com/normalniydada/test_task_infotecs/internal/services"
	"github.com/normalniydada/test_task_infotecs/internal/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetLastTransactions возвращает список последних N транзакций.
//
// GET /api/transactions?count=N
//
// Параметры запроса:
//   - count (int) — количество транзакций для возврата
//
// Ответ:
//   - 200 OK: JSON-массив транзакций
//   - 400 Bad Request: если параметр count некорректный
//   - 500 Internal Server Error: если произошла ошибка при получении данных
func GetLastTransactions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := strconv.Atoi(c.Query("count"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if count <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count value"})
			return
		}

		transactions, err := storage.GetLastNTransactions(db, count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	}
}

// SendTransaction выполняет перевод средств между кошельками
//
// POST /api/send
//
// Тело запроса (JSON):
//
//	{
//	  "from": "wallet1",
//	  "to": "wallet2",
//	  "amount": 33.3
//	}
//
// Поля:
//   - from (string) — адрес отправителя
//   - to (string) — адрес получателя
//   - amount (float64) — сумма перевода в условных единицах (например, 33.3 = 33.3 у.е.)
//
// Ответ:
//   - 200 OK: {"status": "sent"} — если перевод успешен
//   - 400 Bad Request: если входные данные некорректны или недостаточно средств
func SendTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.TransactionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := services.TransferMoney(db, req.From, req.To, convertMoneyToInt(req.Amount))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "sent"})
	}
}

// convertMoneyToInt преобразует сумму в условных единицах (float64) в целочисленное значение (int64)
//
// Например, convertMoneyToInt(50.75) вернёт 5075 (копейки)
func convertMoneyToInt(value float64) int64 {
	return int64(value * 100)
}
