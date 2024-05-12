package city

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, city entity.City) error {
	return u.cityRepository.Create(ctx, &entity.City{
		StateID:   city.StateID,
		Name:      city.Name,
		Latitude:  city.Latitude,
		Longitude: city.Longitude,
		IsActive:  city.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
