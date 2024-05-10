package state

import (
	"github.com/google/uuid"
	"time"
)

type State struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CountryID  uuid.UUID `json:"country_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	IsActive   bool      `json:"is_active"`
	Precedence int64     `json:"precedence"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
