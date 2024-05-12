package country

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, country entity.Country) error {
	return u.countryRepository.Create(ctx, &entity.Country{
		Name:         country.Name,
		Code:         country.Code,
		PhoneCode:    country.PhoneCode,
		Capital:      country.Capital,
		Latitude:     country.Latitude,
		Longitude:    country.Longitude,
		CurrencyCode: country.CurrencyCode,
		IsActive:     country.IsActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
}
