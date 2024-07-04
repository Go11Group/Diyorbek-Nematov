package models

type RegisterReq struct {
	ID   string `json:"user_id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type GetProfileReq struct {
	Email string `json:"email"`
}

type GetProfileResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetProfileByIdReq struct {
	ID string `json:"id"`
}

type GetProfileByIdResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
