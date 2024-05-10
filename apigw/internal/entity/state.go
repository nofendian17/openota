package entity

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type State struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:char(36);not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	CountryID  uuid.UUID `json:"country_id" gorm:"type:char(36);not null"`
	Country    Country   `json:"country" gorm:"foreignKey:CountryID"`
	Cities     []City    `json:"cities" gorm:"foreignKey:StateID"`
	Latitude   float64   `json:"latitude" gorm:"type:DOUBLE PRECISION;null"`
	Longitude  float64   `json:"longitude" gorm:"type:DOUBLE PRECISION;null"`
	IsActive   *bool     `json:"is_active" gorm:"type:boolean;default:true"`
	Precedence int64     `json:"precedence" gorm:"type:integer;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

func (state *State) BeforeCreate(tx *gorm.DB) error {
	state.ID = uuid.New()
	return nil
}
