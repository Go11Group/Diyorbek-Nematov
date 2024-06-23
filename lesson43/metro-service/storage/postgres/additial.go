package postgres

import (
	"database/sql"
	"fmt"
	"metro-service/models"
)

type AdditionalRepo struct {
	DB *sql.DB
}

func NewAddtionalRepo(db *sql.DB) *AdditionalRepo {
	return &AdditionalRepo{DB: db}
}

func (a *AdditionalRepo) GetBalance(userId string) (models.Balance, error) {
	var balance = models.Balance{UserID: userId}
	err := a.DB.QueryRow(`
		SELECT
			COALESCE(ROUND(
				SUM(CASE WHEN t.type = 'debit' THEN t.amount ELSE 0 END)::DECIMAL -
				SUM(CASE WHEN t.type = 'credit' THEN t.amount ELSE 0 END)::DECIMAL, 2
			), 0) AS balance
		FROM
			transactions t
		JOIN
			cards c ON t.card_id = c.id
		WHERE
    		c.user_id = $1
	`, userId).Scan(&balance.Amount)

	return balance, err
}

func (a *AdditionalRepo) CheckBalance(userId string) (models.BalanceResponse, error) {
	var result models.BalanceResponse
	err := a.DB.QueryRow(`
		SELECT
			COALESCE(ROUND(
				SUM(CASE WHEN t.type = 'debit' THEN t.amount ELSE 0 END)::DECIMAL -
				SUM(CASE WHEN t.type = 'credit' THEN t.amount ELSE 0 END)::DECIMAL, 2
			), 0) AS balance,
			CASE
				WHEN COALESCE(ROUND(
					SUM(CASE WHEN t.type = 'debit' THEN t.amount ELSE 0 END)::DECIMAL -
					SUM(CASE WHEN t.type = 'credit' THEN t.amount ELSE 0 END)::DECIMAL, 2
				), 0
				) >= 1700.00 THEN 'Sufficient'
				ELSE 'Insufficient'
			END AS balance_status
		FROM
			transactions t
		INNER JOIN
			cards c ON t.card_id = c.id
		WHERE
			c.user_id = $1
	`, userId).Scan(&result.Balance, &result.BalanceStatus)

	fmt.Println(result, err)

	return result, err
}