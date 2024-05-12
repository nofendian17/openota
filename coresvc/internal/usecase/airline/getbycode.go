package airline

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetByCode(ctx context.Context, code string) (*entity.Airline, error) {
	result, err := u.airlineRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &entity.Airline{
		ID:         result.ID,
		Code:       result.Code,
		Name:       result.Name,
		Logo:       result.Logo,
		IsActive:   result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
