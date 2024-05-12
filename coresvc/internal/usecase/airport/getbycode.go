package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByCode(ctx context.Context, code string) (*entity.Airport, error) {
	result, err := u.airportRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &entity.Airport{
		ID:         result.ID,
		Code:       result.Code,
		Name:       result.Name,
		CityID:     result.CityID,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsDomestic: result.IsDomestic,
		Timezone:   result.Timezone,
		IsActive:   result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
