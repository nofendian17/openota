package city

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/city"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/city"
	cityRepository "github.com/nofendian17/openota/apigw/internal/repository/city"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*response.City, error)
	GetByStateID(ctx context.Context, stateID string) ([]*response.City, error)
	GetAll(ctx context.Context) ([]*response.City, error)

	Create(ctx context.Context, city request.Create) error
	Update(ctx context.Context, ID string, city request.Update) error
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
