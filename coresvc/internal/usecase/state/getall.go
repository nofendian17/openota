package state

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetAll(ctx context.Context) ([]*entity.State, error) {
	result, err := u.stateRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.State, len(result))
	for i, r := range result {
		resp[i] = &entity.State{
			ID:         r.ID,
			Name:       r.Name,
			CountryID:  r.CountryID,
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
