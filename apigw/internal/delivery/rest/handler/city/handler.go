package city

import (
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/gopkg/validator"
	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/usecase/city"
	"net/http"
)

type Handler interface {
	GetByID() http.HandlerFunc
	GetByStateID() http.HandlerFunc
	GetAll() http.HandlerFunc

	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
}

type handler struct {
	config    *config.Config
	useCase   city.UseCase
	logger    logger.Logger
	validator *validator.CustomValidator
}

func New(c *container.Container) Handler {
	return &handler{
		config:    c.Config,
		useCase:   c.UseCase.City,
		logger:    c.Logger,
		validator: c.Validator,
	}
}
