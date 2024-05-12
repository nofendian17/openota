package state

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByCountryID(ctx context.Context, countryID string) ([]*entity.State, error) {
	result, err := u.stateRepository.GetByCountryID(ctx, countryID)
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
