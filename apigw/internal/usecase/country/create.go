package country

import (
	"context"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/country"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, country request.Create) error {
	err := u.countryRepository.Create(ctx, entity.Country{
		Name:         country.Name,
		Code:         country.Code,
		PhoneCode:    country.PhoneCode,
		Capital:      country.Capital,
		Latitude:     country.Latitude,
		Longitude:    country.Longitude,
		CurrencyCode: country.CurrencyCode,
		IsActive:     *country.IsActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
