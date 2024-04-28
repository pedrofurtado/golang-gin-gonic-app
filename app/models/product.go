package models

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID   `json:"id,omitempty" gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Price       float64     `json:"price,omitempty"`
	Quantity    int         `json:"quantity,omitempty"`
	Active      bool        `json:"active,omitempty"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}
