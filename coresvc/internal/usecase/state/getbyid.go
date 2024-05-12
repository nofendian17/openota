package state

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*entity.State, error) {
	result, err := u.stateRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &entity.State{
		ID:         result.ID,
		Name:       result.Name,
		CountryID:  result.CountryID,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
