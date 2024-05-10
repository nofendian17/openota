package state

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/state"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Update(ctx context.Context, ID string, state request.Update) error {
	return u.stateRepository.Update(ctx, ID, &entity.State{
		Name:       state.Name,
		CountryID:  uuid.MustParse(state.CountryID),
		Latitude:   state.Latitude,
		Longitude:  state.Longitude,
		IsActive:   state.IsActive,
		Precedence: state.Precedence,
		UpdatedAt:  time.Now(),
	})
}
