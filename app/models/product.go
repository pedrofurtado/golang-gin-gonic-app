package models

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Active      bool      `json:"active"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
}
