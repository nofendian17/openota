package airport

import (
	"github.com/google/uuid"
	"time"
)

type Airport struct {
	ID         uuid.UUID `json:"id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	CityID     uuid.UUID `json:"city_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	IsDomestic bool      `json:"is_domestic"`
	Timezone   string    `json:"timezone"`
	IsActive   bool      `json:"is_active"`
	Precedence int64     `json:"precedence"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
