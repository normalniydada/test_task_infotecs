package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/normalniydada/test_task_infotecs/internal/services"
	"gorm.io/gorm"
	"net/http"
)

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

func convertMoneyToFloat(value int64) float64 {
	return float64(value) / 100
}
