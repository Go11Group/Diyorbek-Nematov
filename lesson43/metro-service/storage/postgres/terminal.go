package postgres

import (
	"database/sql"
	"metro-service/models"
)

type TerminalRepo struct {
	DB *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{DB: db}
}

func (repo *TerminalRepo) CreateTerminal(terminal models.Terminal) error {
	_, err := repo.DB.Exec(`
		INSERT INTO terminal (
			id, 
			station_id
		) 
			VALUES ($1, $2)
	`, terminal.ID, terminal.StationID)
	return err
}

func (repo *TerminalRepo) GetTerminalByID(id string) (*models.Terminal, error) {
	row := repo.DB.QueryRow(`
		SELECT 
			id,
			station_id 
		FROM 
			terminal 
		WHERE id = $1
	`, id)

	var terminal models.Terminal
	err := row.Scan(&terminal.ID, &terminal.StationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &terminal, nil
}

func (repo *TerminalRepo) GetTerminals() ([]models.Terminal, error) {
	rows, err := repo.DB.Query(`
		SELECT 
			id, 
			station_id 
		FROM terminal
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		if err := rows.Scan(&terminal.ID, &terminal.StationID); err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return terminals, nil
}

func (repo *TerminalRepo) UpdateTerminal(terminal models.Terminal) error {
	query := `UPDATE terminal SET station_id = $1 WHERE id = $2`
	_, err := repo.DB.Exec(query, terminal.StationID, terminal.ID)
	return err
}

func (repo *TerminalRepo) DeleteTerminal(id string) error {
	query := `DELETE FROM terminal WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	return err
}
