package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `jso:"age"`
}

type Card struct {
	ID     string `json:"card_id"`
	UserID string `json:"user_id"`
	Number string `json:"number"`
}

type Transaction struct {
	ID         string  `json:"id"`
	CardID     string  `json:"card_id"`
	Amount     float64 `json:"amount"`
	TerminalID string  `json:"terminal_id"`
	Type       string  `json:"type"`
}

type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Terminal struct {
	ID        string `json:"id"`
	StationID string `json:"station_id"`
}

type Balance struct {
	UserID string `json:"id"`
	Amount float64 `json:"amount"`
}

type BalanceResponse struct {
	Balance      float64 `json:"balance"`
	BalanceStatus string  `json:"balance_status"`
}