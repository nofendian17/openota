package country

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByID(ctx context.Context, ID string) (*entity.Country, error) {
	result, err := u.countryRepository.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &entity.Country{
		ID:           result.ID,
		Name:         result.Name,
		Code:         result.Code,
		PhoneCode:    result.PhoneCode,
		Capital:      result.Capital,
		Latitude:     result.Latitude,
		Longitude:    result.Longitude,
		CurrencyCode: result.CurrencyCode,
		IsActive:     result.IsActive,
		Precedence:   result.Precedence,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}, nil
}
