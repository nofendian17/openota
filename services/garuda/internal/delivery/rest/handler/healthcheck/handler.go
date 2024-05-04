package healthcheck

import (
	"net/http"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/garuda/internal/config"
	"github.com/nofendian17/openota/garuda/internal/container"
	"github.com/nofendian17/openota/garuda/internal/usecase/healthcheck"
)

type Handler interface {
	Ping() http.HandlerFunc
	Readiness() http.HandlerFunc
	Health() http.HandlerFunc
}

type handler struct {
	config  *config.Config
	useCase healthcheck.UseCase
	logger  logger.Logger
}

func New(c *container.Container) Handler {
	return &handler{
		config:  c.Config,
		useCase: c.UseCase.Health,
		logger:  c.Logger,
	}
}
