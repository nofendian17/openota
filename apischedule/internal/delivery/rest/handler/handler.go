package handler

import (
	"github.com/nofendian17/openota/apischedule/internal/container"
	"github.com/nofendian17/openota/apischedule/internal/delivery/rest/handler/healthcheck"
)

type Handler struct {
	Health healthcheck.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health: healthcheck.New(c),
	}
}
