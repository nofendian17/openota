package airline

import (
	"github.com/google/uuid"
	"time"
)

type Airline struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Logo       string    `json:"logo"`
	IsActive   bool      `json:"is_active"`
	Precedence int64     `json:"precedence"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
