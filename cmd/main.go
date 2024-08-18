package main

import (
	"os"
	"net/http"
	log "log"
	"real-estate-service/internal/middleware"
	"real-estate-service/internal/config"
	database "real-estate-service/internal/database"
	"real-estate-service/internal/logger"
	"real-estate-service/api/generated"
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

    options := generated.ChiServerOptions{
        BaseURL:    "/api/v1",
        BaseRouter: r,
        Middlewares: []generated.MiddlewareFunc{
            middleware.LoggerMiddleware(logger),
        },
        ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
            logger.Error("Request handling error", "error", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        },
    }

    server := &handlers.MyServer{
        Logger: logger,
    }

    handler := generated.HandlerWithOptions(server, options)

    r.Get("/dummyLogin", func(w http.ResponseWriter, r *http.Request) {
        params := &generated.GetDummyLoginParams{
            UserType: generated.UserType(r.URL.Query().Get("user_type")),
        }
        server.GetDummyLogin(w, r, params)
    })

    logger.Info("Starting server on port 8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal("Server failed", "error", err)
    }
}