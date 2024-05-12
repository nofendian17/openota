package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetAll(ctx context.Context) ([]*entity.Airport, error) {
	result, err := u.airportRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.Airport, len(result))
	for i, r := range result {
		resp[i] = &entity.Airport{
			ID:         r.ID,
			Code:       r.Code,
			Name:       r.Name,
			CityID:     r.CityID,
			Latitude:   r.Latitude,
			Longitude:  r.Longitude,
			IsDomestic: r.IsDomestic,
			Timezone:   r.Timezone,
			IsActive:   r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
