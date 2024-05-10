package state

import (
	"context"
	"github.com/google/uuid"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/state"
	"github.com/nofendian17/openota/apigw/internal/entity"
	"time"
)

func (u *useCase) Create(ctx context.Context, state request.Create) error {
	return u.stateRepository.Create(ctx, &entity.State{
		Name:      state.Name,
		CountryID: uuid.MustParse(state.CountryID),
		Latitude:  state.Latitude,
		Longitude: state.Longitude,
		IsActive:  state.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
