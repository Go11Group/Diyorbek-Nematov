package handler

import (
	"auth-service/storage/postgres"
)

type Handler struct {
	User postgres.UserRepo
}