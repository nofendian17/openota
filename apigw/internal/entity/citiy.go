package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type City struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:char(36);not null"`
	StateID    uuid.UUID `json:"state_id" gorm:"type:char(36)"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	State      State     `json:"state" gorm:"foreignKey:StateID"`
	Latitude   float64   `json:"latitude" gorm:"type:DOUBLE PRECISION;null"`
	Longitude  float64   `json:"longitude" gorm:"type:DOUBLE PRECISION;null"`
	IsActive   *bool     `json:"is_active" gorm:"type:boolean;default:true"`
	Precedence int64     `json:"precedence" gorm:"type:integer;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

func (city *City) BeforeCreate(tx *gorm.DB) error {
	city.ID = uuid.New()
	return nil
}
