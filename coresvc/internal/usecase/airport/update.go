package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, airport entity.Airport) error {
	return u.airportRepository.Update(ctx, ID, &entity.Airport{
		Code:       airport.Code,
		Name:       airport.Name,
		CityID:     airport.CityID,
		Latitude:   airport.Latitude,
		Longitude:  airport.Longitude,
		IsDomestic: airport.IsDomestic,
		Timezone:   airport.Timezone,
		IsActive:   airport.IsActive,
		Precedence: airport.Precedence,
		UpdatedAt:  time.Now(),
	})
}
