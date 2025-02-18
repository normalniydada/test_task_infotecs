// Package handlers содержит обработчики HTTP-запросов для работы с кошельками
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/normalniydada/test_task_infotecs/internal/services"
	"gorm.io/gorm"
	"net/http"
)

// GetBalance возвращает баланс указанного кошелька
//
// GET /api/wallet/{address}/balance
//
// Параметры запроса:
//   - address (string) — уникальный адрес кошелька.
//
// Ответ:
//   - 200 OK: {"balance": 100.50} — если кошелек найден, баланс возвращается в формате float64 (у.е)
//   - 500 Internal Server Error: {"error": "wallet not found"} — если кошелек не найден или произошла ошибка
//
// Пример запроса:
//
//	GET /api/wallet/wallet123/balance
//
// Пример ответа (успех):
//
//	{
//	  "balance": 150.75
//	}
//
// Пример ответа (ошибка):
//
//	{
//	  "error": "wallet not found"
//	}
func GetBalance(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Param("address")

		balance, err := services.GetWalletBalance(db, address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"balance": convertMoneyToFloat(balance)})
	}
}

// convertMoneyToFloat преобразует баланс из целого числа (копейки) в float64 (у.е)
//
// Например, convertMoneyToFloat(32675) вернёт 326.75
func convertMoneyToFloat(value int64) float64 {
	return float64(value) / 100
}
