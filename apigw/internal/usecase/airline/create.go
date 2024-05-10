package airline

import (
	"context"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airline"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, airline request.Create) error {
	return u.airlineRepository.Create(ctx, &entity.Airline{
		Name:      airline.Name,
		Code:      airline.Code,
		Logo:      airline.Logo,
		IsActive:  airline.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
