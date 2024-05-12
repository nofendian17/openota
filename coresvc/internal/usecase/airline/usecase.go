package airline

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/repository/airline"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*entity.Airline, error)
	GetByCode(ctx context.Context, code string) (*entity.Airline, error)
	GetAll(ctx context.Context) ([]*entity.Airline, error)

	Create(ctx context.Context, airline entity.Airline) error
	Update(ctx context.Context, ID string, airline entity.Airline) error
	Delete(ctx context.Context, ID string) error
}

type useCase struct {
	logger            logger.Logger
	airlineRepository airline.Repository
}

func New(logger logger.Logger, airlineRepository airline.Repository) UseCase {
	return &useCase{
		logger:            logger,
		airlineRepository: airlineRepository,
	}
}
