package airport

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airport"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, airport request.Update) error {
	return u.airportRepository.Update(ctx, ID, &entity.Airport{
		Code:       airport.Code,
		Name:       airport.Name,
		CityID:     uuid.MustParse(airport.CityID),
		Latitude:   airport.Latitude,
		Longitude:  airport.Longitude,
		IsDomestic: airport.IsDomestic,
		Timezone:   airport.Timezone,
		IsActive:   airport.IsActive,
		Precedence: airport.Precedence,
		UpdatedAt:  time.Now(),
	})
}
