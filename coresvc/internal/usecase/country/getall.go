package country

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetAll(ctx context.Context) ([]*entity.Country, error) {
	result, err := u.countryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.Country, len(result))
	for i, r := range result {
		resp[i] = &entity.Country{
			ID:           r.ID,
			Name:         r.Name,
			Code:         r.Code,
			PhoneCode:    r.PhoneCode,
			Capital:      r.Capital,
			Latitude:     r.Latitude,
			Longitude:    r.Longitude,
			CurrencyCode: r.CurrencyCode,
			IsActive:     r.IsActive,
			Precedence:   r.Precedence,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		}
	}
	return resp, nil
}
