package city

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, city request.Update) error {
	return u.cityRepository.Update(ctx, ID, &entity.City{
		StateID:    uuid.MustParse(city.StateID),
		Name:       city.Name,
		Latitude:   city.Latitude,
		Longitude:  city.Longitude,
		IsActive:   city.IsActive,
		Precedence: city.Precedence,
		UpdatedAt:  time.Now(),
	})
}
