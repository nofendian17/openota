package state

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	stateRepository "github.com/nofendian17/openota/coresvc/internal/repository/state"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*entity.State, error)
	GetByCountryID(ctx context.Context, countryID string) ([]*entity.State, error)
	GetAll(ctx context.Context) ([]*entity.State, error)

	Create(ctx context.Context, country entity.State) error
	Update(ctx context.Context, ID string, country entity.State) error
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
