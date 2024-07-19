package models

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Success struct {
	Message string
}

type Errors struct {
	Error string
}
