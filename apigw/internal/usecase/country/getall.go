package country

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/country"
)

func (u *useCase) GetAll(ctx context.Context) ([]*response.Country, error) {
	result, err := u.countryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.Country, len(result))
	for i, r := range result {
		resp[i] = &response.Country{
			ID:           r.ID,
			Name:         r.Name,
			Code:         r.Code,
			Latitude:     r.Latitude,
			Longitude:    r.Longitude,
			CurrencyCode: r.CurrencyCode,
			IsActive:     r.IsActive,
			Precedence:   r.Precedence,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		}
	}
	return resp, nil
}
