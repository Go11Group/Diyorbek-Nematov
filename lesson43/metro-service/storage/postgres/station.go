package postgres

import (
	"metro-service/models"
	"database/sql"
)

type StationRepo struct {
	DB *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{DB: db}
}

// CreateStation yangi stantsiyani qo'shadi
func (repo *StationRepo) CreateStation(station models.Station) error {
	_, err := repo.DB.Exec(`
		INSERT INTO station (
			id, 
			name
		) 
			VALUES ($1, $2)
	`, station.ID, station.Name)
	return err
}

// GetStations barcha stantsiyalarni qaytaradi
func (repo *StationRepo) GetStations() ([]models.Station, error) {
    rows, err := repo.DB.Query(`
		SELECT 
			id, 
			name 
		FROM station
	`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var stations []models.Station
    for rows.Next() {
        var station models.Station
        if err := rows.Scan(&station.ID, &station.Name); err != nil {
            return nil, err
        }
        stations = append(stations, station)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return stations, nil
}

// GetStationByID stantsiyani ID bo'yicha qaytaradi
func (repo *StationRepo) GetStationByID(id string) (*models.Station, error) {
	row := repo.DB.QueryRow(`
		SELECT 
			id, 
			name 
		FROM station 
		WHERE id = $1
	`, id)

	var station models.Station
	err := row.Scan(&station.ID, &station.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &station, nil
}

// UpdateStation stantsiyani yangilaydi
func (repo *StationRepo) UpdateStation(station models.Station) error {
	query := `UPDATE station SET name = $1 WHERE id = $2`
	_, err := repo.DB.Exec(query, station.Name, station.ID)
	return err
}

// DeleteStation stantsiyani ID bo'yicha o'chiradi
func (repo *StationRepo) DeleteStation(id string) error {
	query := `DELETE FROM station WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	return err
}
