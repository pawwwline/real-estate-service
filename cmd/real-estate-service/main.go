package main

import (
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"real-estate-service/internal/db"
	"real-estate-service/internal/middleware"
	"real-estate-service/internal/utils"
	"real-estate-service/repository"

	"net/http"
	"os"
	"real-estate-service/api/generated"
	"real-estate-service/api/handlers"
	"real-estate-service/internal/config"
	"real-estate-service/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	logger := logger.SetupLogger(cfg.Env)
	database, err := db.ConnectDb(&cfg.Storage, logger)
	if err != nil {
		logger.Error("Failed to connect to db", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	err = db.ApplyMigrations(database, logger, &cfg.Storage)
	if err != nil {
		logger.Error("Failed to apply migrations", "error", err)
		os.Exit(1)
	}

	houseRepo := &repository.HouseRepository{Db: database}
	flatRepo := &repository.FlatRepository{Db: database}

	options := generated.ChiServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: chi.NewRouter(),
		Middlewares: []generated.MiddlewareFunc{
			middleware2.RequestID,
			middleware.LoggerMiddleware(logger),
			middleware.TokenAuth,
		},
		ErrorHandlerFunc: utils.ErrorHandlerFunc,
	}

	Myserver := &handlers.MyServer{
		Logger:                   logger,
		HouseRepositoryInterface: houseRepo,
		FlatRepositoryInterface:  flatRepo,
	}

	r := generated.HandlerWithOptions(Myserver, options)

	server := &http.Server{
		Addr:         cfg.HTTPserver.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPserver.Timeout,
		WriteTimeout: cfg.HTTPserver.Timeout,
		IdleTimeout:  cfg.HTTPserver.IdleTimeout,
	}

	logger.Info("Starting server on port 8080")

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server failed", err)
	}
}
