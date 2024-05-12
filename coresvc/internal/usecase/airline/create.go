package airline

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, airline entity.Airline) error {
	return u.airlineRepository.Create(ctx, &entity.Airline{
		Name:      airline.Name,
		Code:      airline.Code,
		Logo:      airline.Logo,
		IsActive:  airline.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
