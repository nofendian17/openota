package state

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, state entity.State) error {
	return u.stateRepository.Create(ctx, &entity.State{
		Name:      state.Name,
		CountryID: state.CountryID,
		Latitude:  state.Latitude,
		Longitude: state.Longitude,
		IsActive:  state.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
