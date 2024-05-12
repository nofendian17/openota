package city

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*entity.City, error) {
	result, err := u.cityRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &entity.City{
		ID:         result.ID,
		StateID:    result.StateID,
		Name:       result.Name,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
