package state

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/state"
)

func (u *useCase) GetByCountryID(ctx context.Context, countryID string) ([]*response.State, error) {
	result, err := u.stateRepository.GetByCountryID(ctx, countryID)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.State, len(result))
	for i, r := range result {
		resp[i] = &response.State{
			ID:         r.ID,
			Name:       r.Name,
			CountryID:  r.CountryID,
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
