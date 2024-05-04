package handler

import (
	"github.com/nofendian17/openota/apibook/internal/container"
	"github.com/nofendian17/openota/apibook/internal/delivery/rest/handler/healthcheck"
)

type Handler struct {
	Health healthcheck.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health: healthcheck.New(c),
	}
}
