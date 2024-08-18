package main

import (
	"os"
	"net/http"
	log "log"
	"real-estate-service/internal/middleware"
	"real-estate-service/internal/config"
	database "real-estate-service/internal/database"
	"real-estate-service/internal/logger"
	"github.com/go-chi/chi/v5"
	_"github.com/golang-migrate/migrate/v4/database/postgres"
	"real-estate-service/api/handlers"


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

	r := chi.NewRouter()

    r.Use(middleware.LoggerMiddleware(logger))

    server := &handlers.MyServer{
        Logger: logger,
    }

    r.Get("/dummyLogin", func(w http.ResponseWriter, r *http.Request) {
        userType := r.URL.Query().Get("user_type")
        params := handlers.GetDummyLoginParams{
            UserType: handlers.UserType(userType),
        }
        server.GetDummyLogin(w, r, params)
    })

    logger.Info("Starting server on port 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Server failed", "error", err)
    }
}
