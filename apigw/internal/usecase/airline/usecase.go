package airline

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airline"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airline"
	"github.com/nofendian17/openota/apigw/internal/repository/airline"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*response.Airline, error)
	GetByCode(ctx context.Context, code string) (*response.Airline, error)
	GetAll(ctx context.Context) ([]*response.Airline, error)

	Create(ctx context.Context, airline request.Create) error
	Update(ctx context.Context, ID string, airline request.Update) error
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
