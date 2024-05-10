package airline

import (
	"context"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airline"
)

func (u *useCase) GetAll(ctx context.Context) ([]*response.Airline, error) {
	result, err := u.airlineRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*response.Airline, len(result))
	for i, r := range result {
		resp[i] = &response.Airline{
			ID:         r.ID,
			Code:       r.Code,
			Name:       r.Name,
			Logo:       r.Logo,
			IsActive:   *r.IsActive,
			Precedence: r.Precedence,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		}
	}
	return resp, nil
}
