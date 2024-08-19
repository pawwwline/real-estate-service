package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"net/http"
	"os"
	"real-estate-service/api/generated"
	"real-estate-service/api/handlers"
	"real-estate-service/internal/config"
	database "real-estate-service/internal/database"
	"real-estate-service/internal/logger"
	"real-estate-service/internal/middleware"
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

	options := generated.ChiServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: chi.NewRouter(),
		Middlewares: []generated.MiddlewareFunc{
			middleware.LoggerMiddleware(logger),
		},
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			logger.Error("Request handling error", "error", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		},
	}

	Myserver := &handlers.MyServer{
		Logger: logger,
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
