package airport

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airport"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, airport request.Create) error {
	return u.airportRepository.Create(ctx, &entity.Airport{
		Code:       airport.Code,
		Name:       airport.Name,
		CityID:     uuid.MustParse(airport.CityID),
		Latitude:   airport.Latitude,
		Longitude:  airport.Longitude,
		IsDomestic: airport.IsDomestic,
		Timezone:   airport.Timezone,
		IsActive:   airport.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
}
