package state

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, state entity.State) error {
	return u.stateRepository.Update(ctx, ID, &entity.State{
		Name:       state.Name,
		CountryID:  state.CountryID,
		Latitude:   state.Latitude,
		Longitude:  state.Longitude,
		IsActive:   state.IsActive,
		Precedence: state.Precedence,
		UpdatedAt:  time.Now(),
	})
}
