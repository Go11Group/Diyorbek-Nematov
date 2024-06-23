package postgres

import (
	"database/sql"
	"errors"
	"metro-service/models"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}

func (repo *TransactionRepo) GetCardBalance(tx *sql.Tx, cardID string) (float64, error) {
	var balance float64
	err := tx.QueryRow(`
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
    		c.id = $1
	`, cardID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (repo *TransactionRepo) CreateTransaction(transaction models.Transaction) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	// Tranzaksiyani yakunlash yoki orqaga qaytarishni nazorat qilish
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Karta balansini olish
	balance, err := repo.GetCardBalance(tx, transaction.CardID)
	if err != nil {
		return err
	}

	// Tranzaksiyani kiritishdan oldin balansni tekshirish
	if transaction.Type == "credit" && balance < transaction.Amount {
		return errors.New("insufficient balance")
	}

	// Tranzaksiyani kiritish
	_, err = tx.Exec(`
		INSERT INTO transactions (
			id, 
			card_id, 
			terminal_id, 
			amount, 
			type
		) 
		VALUES ($1, $2, $3, $4, $5)
	`, transaction.ID, transaction.CardID, transaction.TerminalID, transaction.Amount, transaction.Type)
	if err != nil {
		return err
	}

	// Tranzaksiyani muvaffaqiyatli yakunlash
	return nil
}

func (repo *TransactionRepo) GetTransactionByID(id string) (*models.Transaction, error) {
	row := repo.DB.QueryRow(`
		SELECT 
			id, 
			card_id, 
			terminal_id, 
			amount, 
			type 
		FROM 
			transactions 
		WHERE id = $1
	`, id)

	var transaction models.Transaction
	err := row.Scan(&transaction.ID, &transaction.CardID, &transaction.TerminalID, &transaction.Amount, &transaction.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (repo *TransactionRepo) GetTransactions() ([]models.Transaction, error) {
	rows, err := repo.DB.Query(`
		SELECT 
			id, 
			card_id, 
			terminal_id, 
			amount, 
			type 
		FROM 
			transactions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CardID, &transaction.TerminalID,
			&transaction.Amount, &transaction.Type); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo *TransactionRepo) UpdateTransaction(transaction models.Transaction) error {
	_, err := repo.DB.Exec(`
		UPDATE 
			transactions SET 
			card_id = $1, 
			terminal_id = $2, 
			amount = $3, 
			type = $4 
		WHERE id = $5
	`, transaction.CardID, transaction.TerminalID, transaction.Amount, transaction.Type, transaction.ID)
	return err
}

func (repo *TransactionRepo) DeleteTransaction(id string) error {
	query := `DELETE FROM transactions WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	return err
}
