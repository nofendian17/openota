package handler

import (
	"github.com/nofendian17/openota/apiticketing/internal/container"
	"github.com/nofendian17/openota/apiticketing/internal/delivery/rest/handler/healthcheck"
)

type Handler struct {
	Health healthcheck.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health: healthcheck.New(c),
	}
}
