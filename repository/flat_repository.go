package repository

import (
	"database/sql"
	"fmt"
	"real-estate-service/api/generated"
)

type FlatRepositoryInterface interface {
	CreateFlat(flat *generated.Flat) error
	UpdateFlat(flat *generated.Flat) error
	GetFlatId(flatId generated.FlatId) (*generated.Flat, error)
	GetFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error)
	GetApprovedFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error)
}

type FlatRepository struct {
	Db *sql.DB
}

func (repo *FlatRepository) CreateFlat(flat *generated.Flat) error {

	query := "INSERT INTO flat (house_id, price, rooms) VALUES ($1, $2, $3)"
	_, err := repo.Db.Exec(query, flat.HouseId, flat.Price, flat.Rooms) //Возможно лучше поставить default в бд
	if err != nil {
		return fmt.Errorf("failed to insert flat: %w", err)
	}

	houseQuery := `
		UPDATE house
		SET updated_at = NOW()
		WHERE id = $1`
	_, err = repo.Db.Exec(houseQuery, flat.HouseId)
	if err != nil {
		return fmt.Errorf("failed to update house: %w", err)
	}

	return nil
}

func (repo *FlatRepository) GetFlatId(flatId generated.FlatId) (*generated.Flat, error) {
	query := "SELECT id, status FROM flat WHERE id = $1"
	flat := &generated.Flat{}
	row := repo.Db.QueryRow(query, flatId)
	err := row.Scan(&flat.Id, &flat.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("flat not found")
		}
		return nil, fmt.Errorf("failed to retrieve flat: %w", err)
	}

	return flat, nil

}

func (repo *FlatRepository) UpdateFlat(flat *generated.Flat) error {
	validStatuses := map[string]bool{
		"created":       true,
		"approved":      true,
		"declined":      true,
		"on moderation": true,
	}
	if !validStatuses[string(flat.Status)] {
		return fmt.Errorf("invalid status: %s", flat.Status)
	}

	if flat.Status == "on moderation" {
		var currentStatus string
		query := "SELECT status FROM flat WHERE id = $1"
		err := repo.Db.QueryRow(query, flat.Id).Scan(&currentStatus)
		if err != nil {
			return fmt.Errorf("failed to check current status: %w", err)
		}
		if currentStatus == "on moderation" {
			return fmt.Errorf("flat is already under moderation by another moderator")
		}
	}

	query := "UPDATE flat SET status = $1 WHERE id = $2"
	_, err := repo.Db.Exec(query, flat.Status, flat.Id)
	return err
}

func (repo *FlatRepository) GetFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error) {
	query := "SELECT id, house_id, price, rooms, status FROM flat WHERE house_id = $1"
	rows, err := repo.Db.Query(query, houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []generated.Flat
	for rows.Next() {
		var flat generated.Flat
		if err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}
	return flats, nil
}

func (repo *FlatRepository) GetApprovedFlatsByHouseId(houseId generated.HouseId) ([]generated.Flat, error) {
	query := "SELECT id, house_id, price, rooms, status FROM flat WHERE house_id = $1 AND status = 'approved'"
	rows, err := repo.Db.Query(query, houseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flats []generated.Flat
	for rows.Next() {
		var flat generated.Flat
		if err := rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}
	return flats, nil
}
