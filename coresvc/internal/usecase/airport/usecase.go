package airport

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"github.com/nofendian17/openota/coresvc/internal/repository/airport"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*entity.Airport, error)
	GetByCode(ctx context.Context, code string) (*entity.Airport, error)
	GetAll(ctx context.Context) ([]*entity.Airport, error)

	Create(ctx context.Context, airport entity.Airport) error
	Update(ctx context.Context, ID string, airport entity.Airport) error
	Delete(ctx context.Context, ID string) error
}

type useCase struct {
	logger            logger.Logger
	airportRepository airport.Repository
}

func New(logger logger.Logger, airportRepository airport.Repository) UseCase {
	return &useCase{
		logger:            logger,
		airportRepository: airportRepository,
	}
}
