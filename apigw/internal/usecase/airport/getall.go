package airport

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airport"
)

func (u *useCase) GetAll(ctx context.Context) ([]*response.Airport, error) {
	result, err := u.airportRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.Airport, len(result))
	for i, r := range result {
		resp[i] = &response.Airport{
			ID:         r.ID,
			Code:       r.Code,
			Name:       r.Name,
			CityID:     r.CityID,
			Latitude:   r.Latitude,
			Longitude:  r.Longitude,
			IsDomestic: *r.IsDomestic,
			Timezone:   r.Timezone,
			IsActive:   *r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
