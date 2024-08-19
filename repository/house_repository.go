package repository

import (
	"database/sql"
	"fmt"
	"real-estate-service/api/generated"
)

type HouseRepositoryInterface interface {
	CreateHouse(house *generated.House) error
}
type HouseRepository struct {
	db *sql.DB
}

func (repo *HouseRepository) CreateHouse(house *generated.House) error {
	query := "INSERT INTO houses (address, year, developer) VALUES ($1, $2, $3)"
	_, err := repo.db.Exec(query, house.Address, house.Year, house.Developer)
	if err != nil {
		return fmt.Errorf("failed to insert house: %w", err)
	}
	return nil
}
