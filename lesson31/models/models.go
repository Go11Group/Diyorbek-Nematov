package models

import (
	"time"
)

type Product struct {
	ID string
	Name string
	Price float64
	CreatedAt time.Time
}