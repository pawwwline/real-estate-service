package db

import (
	"database/sql"
	"fmt"
    "path/filepath"
	"log/slog"
	"real-estate-service/internal/config"
	_"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func ConnectDb(cfg *config.Storage, log *slog.Logger) (*sql.DB, error) {
	fmt.Printf("Config values: host=%s, port=%s, user=%s, password=%s, name=%s\n",
	cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	fmt.Printf("Using connection string: %s", connStr)
	db, err := sql.Open("postgres", connStr)
	//TO:DO add logger
	if err != nil {
		log.Error("Failed to open connection to DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)

	}

	return db, nil
}


//TO:DO Add logger
func ApplyMigrations(db *sql.DB, log *slog.Logger, cfg *config.Storage) error  {

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	absPath, err := filepath.Abs("/Users/polinakuznecova/real-estate-service/internal/db/migrations")
	if err != nil {
		log.Error("Failed to get absolute path", "error", err)
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	sourceURL := fmt.Sprintf("file://%s", absPath)
	m, err := migrate.New(
		sourceURL,
		connStr,
	)
	if err != nil {
		log.Error("failed to create migrate instance", "error", err)
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed to apply migrations", "error", err)
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
    
	return nil
}