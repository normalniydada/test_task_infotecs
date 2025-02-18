package main

import (
	"github.com/gin-gonic/gin"
	"github.com/normalniydada/test_task_infotecs/internal/config"
	"github.com/normalniydada/test_task_infotecs/internal/handlers"
	"github.com/normalniydada/test_task_infotecs/internal/seeds"
	"github.com/normalniydada/test_task_infotecs/internal/storage"
	"github.com/normalniydada/test_task_infotecs/pkg/logger"
	"go.uber.org/zap"
)

// TODO: docker
// TODO: git

func main() {
	// init logger
	zLog := logger.InitLogger()
	defer zLog.Sync()

	// Off log gin
	gin.SetMode(gin.ReleaseMode)

	// config
	cfg := config.MustLoad(zLog)

	// connection db
	db := storage.InitDB(&cfg.Database, zLog)
	defer storage.CloseDB(db, zLog)

	// init 10 wallets
	seeds.InitWallets(db, zLog)

	// create server
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/send", handlers.SendTransaction(db))
		api.GET("/transactions", handlers.GetLastTransactions(db))
		api.GET("/wallet/:address/balance", handlers.GetBalance(db))
	}

	zLog.Info("Server is running...", zap.String("address", cfg.Server.Address))
	if err := r.Run(cfg.Server.Address); err != nil {
		zLog.Fatal("Error start the server", zap.Error(err))
	}
}
