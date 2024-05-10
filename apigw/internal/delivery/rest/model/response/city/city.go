package city

import (
	"github.com/google/uuid"
	"time"
)

type City struct {
	ID         uuid.UUID `json:"id"`
	StateID    uuid.UUID `json:"state_id"`
	Name       string    `json:"name"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	IsActive   bool      `json:"is_active"`
	Precedence int64     `json:"precedence"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
