package country

import (
	"context"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/country"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, country request.Update) error {
	err := u.countryRepository.Update(ctx, ID, entity.Country{
		Name:         country.Name,
		Code:         country.Code,
		PhoneCode:    country.PhoneCode,
		Capital:      country.Capital,
		Latitude:     country.Latitude,
		Longitude:    country.Longitude,
		CurrencyCode: country.CurrencyCode,
		IsActive:     country.IsActive,
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
