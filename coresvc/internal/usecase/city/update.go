package city

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, city entity.City) error {
	return u.cityRepository.Update(ctx, ID, &entity.City{
		StateID:    city.StateID,
		Name:       city.Name,
		Latitude:   city.Latitude,
		Longitude:  city.Longitude,
		IsActive:   city.IsActive,
		Precedence: city.Precedence,
		UpdatedAt:  time.Now(),
	})
}
