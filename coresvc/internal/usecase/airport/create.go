package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, airport entity.Airport) error {
	return u.airportRepository.Create(ctx, &entity.Airport{
		Code:       airport.Code,
		Name:       airport.Name,
		CityID:     airport.CityID,
		Latitude:   airport.Latitude,
		Longitude:  airport.Longitude,
		IsDomestic: airport.IsDomestic,
		Timezone:   airport.Timezone,
		IsActive:   airport.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
}
