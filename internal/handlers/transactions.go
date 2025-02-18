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

func convertMoneyToInt(value float64) int64 {
	return int64(value * 100)
}
