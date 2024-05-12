package country

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	countryRepository "github.com/nofendian17/openota/coresvc/internal/repository/country"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*entity.Country, error)
	GetByCode(ctx context.Context, code string) (*entity.Country, error)
	GetAll(ctx context.Context) ([]*entity.Country, error)

	Create(ctx context.Context, country entity.Country) error
	Update(ctx context.Context, ID string, country entity.Country) error
	Delete(ctx context.Context, ID string) error
}

type useCase struct {
	logger            logger.Logger
	countryRepository countryRepository.Repository
}

func New(logger logger.Logger, countryRepository countryRepository.Repository) UseCase {
	return &useCase{
		logger:            logger,
		countryRepository: countryRepository,
	}
}
