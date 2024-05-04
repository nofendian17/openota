package handler

import (
	"github.com/nofendian17/openota/apifare/internal/container"
	"github.com/nofendian17/openota/apifare/internal/delivery/rest/handler/healthcheck"
)

type Handler struct {
	Health healthcheck.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health: healthcheck.New(c),
	}
}
