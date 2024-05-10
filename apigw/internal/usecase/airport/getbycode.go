package airport

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airport"
)

func (u *useCase) GetByCode(ctx context.Context, code string) (*response.Airport, error) {
	result, err := u.airportRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &response.Airport{
		ID:         result.ID,
		Code:       result.Code,
		Name:       result.Name,
		CityID:     result.CityID,
		Latitude:   result.Latitude,
		Longitude:  result.Longitude,
		IsDomestic: *result.IsDomestic,
		Timezone:   result.Timezone,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
