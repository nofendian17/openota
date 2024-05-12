package city

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetAll(ctx context.Context) ([]*entity.City, error) {
	result, err := u.cityRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.City, len(result))
	for i, r := range result {
		resp[i] = &entity.City{
			ID:         r.ID,
			Name:       r.Name,
			StateID:    r.StateID,
			Latitude:   r.Latitude,
			Longitude:  r.Longitude,
			IsActive:   r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
