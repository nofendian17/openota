package country

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, country entity.Country) error {
	return u.countryRepository.Update(ctx, ID, &entity.Country{
		Name:         country.Name,
		Code:         country.Code,
		PhoneCode:    country.PhoneCode,
		Capital:      country.Capital,
		Latitude:     country.Latitude,
		Longitude:    country.Longitude,
		CurrencyCode: country.CurrencyCode,
		IsActive:     country.IsActive,
		Precedence:   country.Precedence,
		UpdatedAt:    time.Now(),
	})

}
