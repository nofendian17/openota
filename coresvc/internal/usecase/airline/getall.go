package airline

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
)

func (u *useCase) GetAll(ctx context.Context) ([]*entity.Airline, error) {
	result, err := u.airlineRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*entity.Airline, len(result))
	for i, r := range result {
		resp[i] = &entity.Airline{
			ID:         r.ID,
			Code:       r.Code,
			Name:       r.Name,
			Logo:       r.Logo,
			IsActive:   r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
