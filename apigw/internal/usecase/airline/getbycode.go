package airline

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airline"
)

func (u *useCase) GetByCode(ctx context.Context, code string) (*response.Airline, error) {
	result, err := u.airlineRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &response.Airline{
		ID:         result.ID,
		Code:       result.Code,
		Name:       result.Name,
		Logo:       result.Logo,
		IsActive:   *result.IsActive,
		Precedence: result.Precedence,
		CreatedAt:  result.CreatedAt,
		UpdatedAt:  result.UpdatedAt,
	}, nil
}
