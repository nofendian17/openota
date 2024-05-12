package airport

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*entity.Airport, error) {
	result, err := u.airportRepository.GetByID(ctx, ID)
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
