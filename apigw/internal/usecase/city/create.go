package city

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, city request.Create) error {
	return u.cityRepository.Create(ctx, &entity.City{
		StateID:   uuid.MustParse(city.StateID),
		Name:      city.Name,
		Latitude:  city.Latitude,
		Longitude: city.Longitude,
		IsActive:  city.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
