package handler

import (
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/country"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/healthcheck"
)

type Handler struct {
	Health  healthcheck.Handler
	Country country.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health:  healthcheck.New(c),
		Country: country.New(c),
	}
}
