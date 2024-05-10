package airport

import (
	"context"
	"github.com/nofendian17/gopkg/logger"
	request "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/request/airport"
	response "github.com/nofendian17/openota/apigw/internal/delivery/rest/model/response/airport"
	"github.com/nofendian17/openota/apigw/internal/repository/airport"
)

type UseCase interface {
	GetByID(ctx context.Context, ID string) (*response.Airport, error)
	GetByCode(ctx context.Context, code string) (*response.Airport, error)
	GetAll(ctx context.Context) ([]*response.Airport, error)

	Create(ctx context.Context, airport request.Create) error
	Update(ctx context.Context, ID string, airport request.Update) error
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
