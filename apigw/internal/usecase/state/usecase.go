package state

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/state"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/state"
	stateRepository "github.com/nofendian17/openota/apigw/internal/repository/state"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*response.State, error)
	GetByCountryID(ctx context.Context, countryID string) ([]*response.State, error)
	GetAll(ctx context.Context) ([]*response.State, error)

	Create(ctx context.Context, country request.Create) error
	Update(ctx context.Context, ID string, country request.Update) error
	Delete(ctx context.Context, ID string) error
}

type useCase struct {
	logger          logger.Logger
	stateRepository stateRepository.Repository
}

func New(logger logger.Logger, stateRepository stateRepository.Repository) UseCase {
	return &useCase{
		logger:          logger,
		stateRepository: stateRepository,
	}
}
