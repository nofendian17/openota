package airline

import (
	"context"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airline"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, airline request.Update) error {
	return u.airlineRepository.Update(ctx, ID, &entity.Airline{
		Name:       airline.Name,
		Code:       airline.Code,
		Logo:       airline.Logo,
		IsActive:   airline.IsActive,
		Precedence: airline.Precedence,
		UpdatedAt:  time.Now(),
	})
}
