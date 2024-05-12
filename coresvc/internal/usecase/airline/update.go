package airline

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, airline entity.Airline) error {
	return u.airlineRepository.Update(ctx, ID, &entity.Airline{
		Name:       airline.Name,
		Code:       airline.Code,
		Logo:       airline.Logo,
		IsActive:   airline.IsActive,
		Precedence: airline.Precedence,
		UpdatedAt:  time.Now(),
	})
}
