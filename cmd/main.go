package main

import (
	"os"
	"real-estate-service/internal/config"
	"real-estate-service/internal/database"
	"real-estate-service/internal/utils/logger"
	_"github.com/golang-migrate/migrate/v4/database/postgres"

)

func main() {
	cfg, err := config.LoadConfig()
	logger := logger.SetupLogger(cfg.Env)
	db, err := database.ConnectDb(&cfg.Storage, logger)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	
	err = database.ApplyMigrations(db, logger, &cfg.Storage)
	if err != nil {
		logger.Error("Failed to apply migrations", "error", err)
		os.Exit(1)
	}

	logger.Info("Application started successfully")

}
