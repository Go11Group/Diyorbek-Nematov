package models

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
