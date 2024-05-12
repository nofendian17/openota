package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Country struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:char(36);not null"`
	Name         string    `json:"name" gorm:"type:varchar(255);not null;unique"`
	States       []State   `json:"states" gorm:"foreignKey:CountryID;references:ID"`
	Code         string    `json:"code" gorm:"type:varchar(3);unique;not null"`
	PhoneCode    string    `json:"phone_code" gorm:"type:varchar(5);not null"`
	Capital      string    `json:"capital" gorm:"type:varchar(255);not null"`
	Latitude     float64   `json:"latitude" gorm:"type:DOUBLE PRECISION;null"`
	Longitude    float64   `json:"longitude" gorm:"type:DOUBLE PRECISION;null"`
	CurrencyCode string    `json:"currency" gorm:"type:varchar(3);not null"`
	IsActive     *bool     `json:"is_active" gorm:"type:boolean;default:false"`
	Precedence   int64     `json:"precedence" gorm:"type:integer;default:0"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

func (country *Country) BeforeCreate(tx *gorm.DB) error {
	country.ID = uuid.New()
	return nil
}
