package storage

import (
	"fmt"
	"github.com/normalniydada/test_task_infotecs/internal/config"
	"github.com/normalniydada/test_task_infotecs/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.DatabaseConfig, zLog *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSlMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zLog.Fatal("Database connection error: ", zap.Error(err))
	}

	zLog.Info("Database connection success",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
	)

	if err = db.AutoMigrate(&models.Wallet{}, &models.Transaction{}); err != nil {
		zLog.Fatal("Database migration error: ", zap.Error(err))
	}
	zLog.Info("Database migration success")

	return db
}

func CloseDB(db *gorm.DB, zLog *zap.Logger) {
	pdb, err := db.DB()
	if err != nil {
		zLog.Fatal("Error getting database: ", zap.Error(err))
		return
	}

	if err = pdb.Close(); err != nil {
		zLog.Fatal("Error closing database: ", zap.Error(err))
	} else {
		zLog.Info("Closing database success")
	}
}

/*func Truncate(db *gorm.DB) error {
	if err := db.Exec("TRUNCATE TABLE wallets, transactions RESTART IDENTITY CASCADE").Error; err != nil {
		return err
	}
	log.Println("[DB] Tables clear!")
	return nil
}*/
