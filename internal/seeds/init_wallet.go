package seeds

import (
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitWallets(db *gorm.DB, zLog *zap.Logger) {
	var countWallets int64
	db.Model(&models.Wallet{}).Count(&countWallets)

	if countWallets > 0 {
		return
	}

	wallets := make([]models.Wallet, 10)
	for i := 0; i < 10; i++ {
		wallet := models.Wallet{Balance: 10000}
		wallet.CreateWalletAddress()
		wallets[i] = wallet
	}

	if err := db.Create(&wallets).Error; err != nil {
		zLog.Fatal("Error init wallet: ", zap.Error(err))
	}

	zLog.Info("Init wallets successfully")
}
