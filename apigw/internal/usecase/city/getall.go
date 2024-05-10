package city

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/city"
)

func (u *useCase) GetAll(ctx context.Context) ([]*response.City, error) {
	result, err := u.cityRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.City, len(result))
	for i, r := range result {
		resp[i] = &response.City{
			ID:         r.ID,
			Name:       r.Name,
			Latitude:   r.Latitude,
			Longitude:  r.Longitude,
			IsActive:   *r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
