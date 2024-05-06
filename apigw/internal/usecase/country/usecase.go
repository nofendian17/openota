package country

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/country"
	countryRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*country.Country, error)
	GetByCode(ctx context.Context, Code string) (*country.Country, error)
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
