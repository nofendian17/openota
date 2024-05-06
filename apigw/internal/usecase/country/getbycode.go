package country

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/country"
)

func (u *useCase) GetByCode(ctx context.Context, Code string) (*response.Country, error) {
	result, err := u.countryRepository.GetByCode(ctx, Code)
	if err != nil {
		return nil, err
	}

	return &response.Country{
		ID:           result.ID,
		Name:         result.Name,
		Code:         result.Code,
		PhoneCode:    result.PhoneCode,
		Capital:      result.Capital,
		Latitude:     result.Latitude,
		Longitude:    result.Longitude,
		CurrencyCode: result.CurrencyCode,
		IsActive:     result.IsActive,
		Precedence:   result.Precedence,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}, nil
}
