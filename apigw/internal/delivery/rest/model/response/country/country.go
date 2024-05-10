package country

import (
	"github.com/google/uuid"
	"time"
)

type Country struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	PhoneCode    string    `json:"phone_code" `
	Capital      string    `json:"capital"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	CurrencyCode string    `json:"currency"`
	IsActive     bool      `json:"is_active"`
	Precedence   int64     `json:"precedence"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
