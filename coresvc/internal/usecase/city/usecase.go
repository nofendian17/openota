package city

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	cityRepository "github.com/nofendian17/openota/coresvc/internal/repository/city"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*entity.City, error)
	GetByStateID(ctx context.Context, stateID string) ([]*entity.City, error)
	GetAll(ctx context.Context) ([]*entity.City, error)

	Create(ctx context.Context, city entity.City) error
	Update(ctx context.Context, ID string, city entity.City) error
	Delete(ctx context.Context, ID string) error
}

type useCase struct {
	logger         logger.Logger
	cityRepository cityRepository.Repository
}

func New(logger logger.Logger, cityRepository cityRepository.Repository) UseCase {
	return &useCase{
		logger:         logger,
		cityRepository: cityRepository,
	}
}
