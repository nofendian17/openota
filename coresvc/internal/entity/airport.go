package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Airport struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:char(36);not null"`
	Code       string    `json:"code" gorm:"type:varchar(3);unique_index;not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	CityID     uuid.UUID `json:"city_id" gorm:"type:char(36);not null"`
	City       City      `json:"city" gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Latitude   float64   `json:"latitude" gorm:"type:DOUBLE PRECISION;null"`
	Longitude  float64   `json:"longitude" gorm:"type:DOUBLE PRECISION;null"`
	IsDomestic *bool     `json:"is_domestic" gorm:"type:boolean;default:false"`
	Timezone   string    `json:"timezone" gorm:"type:varchar(50);not null"`
	IsActive   *bool     `json:"is_active" gorm:"type:boolean;default:true"`
	Precedence int64     `json:"precedence" gorm:"type:integer;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}

func (airport *Airport) BeforeCreate(tx *gorm.DB) error {
	airport.ID = uuid.New()
	return nil
}
