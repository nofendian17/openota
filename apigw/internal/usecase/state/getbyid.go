package state

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/state"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*response.State, error) {
	result, err := u.stateRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &response.State{
		ID:         result.ID,
		Name:       result.Name,
		CountryID:  result.CountryID,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
