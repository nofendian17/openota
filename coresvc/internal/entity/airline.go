package entity

import (
	"github.com/google/uuid"
	"time"
)

type Airline struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:char(36);not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	Code       string    `json:"code" gorm:"type:varchar(2);unique_index;not null"`
	Logo       string    `json:"logo" gorm:"type:varchar(255)"`
	IsActive   *bool     `json:"is_active" gorm:"type:boolean;default:true"`
	Precedence int64     `json:"precedence" gorm:"type:integer;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
