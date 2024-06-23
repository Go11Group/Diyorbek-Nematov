package postgres

import (
	"database/sql"
	"metro-service/models"
	"metro-service/pkg"
)

type CardRepo struct {
	DB *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{DB: db}
}

func (c *CardRepo) CreateCard(card models.Card) error {
	_, err := c.DB.Exec(`
		INSERT INTO cards (
			id, 
			user_id, 
			number
		)
			VALUES($1, $2)
	`, card.ID, card.UserID, card.Number)

	return err
}

func (c *CardRepo) GetCards() ([]models.Card, error) {
	var cards []models.Card

	rows, err := c.DB.Query(`
		SELECT
			id,
			user_id,
			number
		FROM
			cards
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var card models.Card

		err = rows.Scan(&card.ID, &card.UserID, &card.Number)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepo) GetCardByID(id string) (models.Card, error) {
	var card models.Card

	err := c.DB.QueryRow(`
		SELECT
			id,
			user_id,
			number
		FROM
			cards
		WHERE 
			id = $1
	`, id).Scan(&card.ID, &card.UserID, &card.Number)

	return card, err
}

func (c *CardRepo) UpdateCard(card models.Card) error {
	params := make(map[string]interface{})
	var query = "UPDATE cards SET "
	if card.UserID != "" {
		query += "user_id = :user_id, "
		params["user_id"] = card.UserID
	}
	if card.Number != "" {
		query += "number = :number, "
		params["number"] = card.Number
	}

	query += "WHERE course_id = :id"
	params["id"] = card.ID

	query, args := pkg.ReplaceQueryParams(query, params)

	_, err := c.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CardRepo) DeleteCard(id string) error {
    query := "DELETE FROM cards WHERE id = $1"
    _, err := c.DB.Exec(query, id)
    if err != nil {
        return err
    }
    return nil
}
