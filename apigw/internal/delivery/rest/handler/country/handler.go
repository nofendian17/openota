package country

import (
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/gopkg/validator"
	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/usecase/country"
	"net/http"
)

type Handler interface {
	GetByID() http.HandlerFunc
	GetByCode() http.HandlerFunc
	GetAll() http.HandlerFunc

	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
}

type handler struct {
	config    *config.Config
	useCase   country.UseCase
	logger    logger.Logger
	validator *validator.CustomValidator
}

func New(c *container.Container) Handler {
	return &handler{
		config:    c.Config,
		useCase:   c.UseCase.Country,
		logger:    c.Logger,
		validator: c.Validator,
	}
}