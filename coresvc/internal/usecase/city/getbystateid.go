package city

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByStateID(ctx context.Context, stateID string) ([]*entity.City, error) {
	result, err := u.cityRepository.GetByStateID(ctx, stateID)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.City, len(result))
	for i, r := range result {
		resp[i] = &entity.City{
			ID:         r.ID,
			StateID:    r.StateID,
			Name:       r.Name,
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
