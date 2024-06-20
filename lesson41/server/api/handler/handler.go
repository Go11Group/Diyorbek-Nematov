package handler

import "server/storage/postgres"

type Handler struct {
	Emp postgres.EmployeeRepo
}
