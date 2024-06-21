package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}


type FilterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type SearchUser struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	AgeFrom int    `json:"age_from"`
	AgeTo   int    `json:"age_to"`
}


type AdditialUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

