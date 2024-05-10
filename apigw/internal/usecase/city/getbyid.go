package city

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/city"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*response.City, error) {
	result, err := u.cityRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &response.City{
		ID:         result.ID,
		StateID:    result.StateID,
		Name:       result.Name,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
