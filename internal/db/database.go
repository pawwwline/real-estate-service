package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log/slog"
	"path/filepath"
	"real-estate-service/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDb(cfg *config.Storage, log *slog.Logger) (*sql.DB, error) {
	log.Debug("Config values",
		"host", cfg.DbHost,
		"port", cfg.DbPort,
		"user", cfg.DbUser,
		"password", cfg.DbPassword,
		"name", cfg.DbName,
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	log.Debug("Using connection string:", connStr)
	db, err := sql.Open("postgres", connStr)
	//TO:DO add logger
	if err != nil {
		log.Error("Failed to open connection to DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)

	}

	return db, nil
}

// ApplyMigrations TO:DO Add logger
func ApplyMigrations(db *sql.DB, log *slog.Logger, cfg *config.Storage) error {

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

	if err := m.Up(); err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Error("failed to apply migrations", "error", err)
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}
