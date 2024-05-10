package handler

import (
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/airport"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/city"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/country"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/healthcheck"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/state"
)

type Handler struct {
	Health  healthcheck.Handler
	Country country.Handler
	State   state.Handler
	City    city.Handler
	Airport airport.Handler
}

func New(c *container.Container) *Handler {
	return &Handler{
		Health:  healthcheck.New(c),
		Country: country.New(c),
		State:   state.New(c),
		City:    city.New(c),
		Airport: airport.New(c),
	}
}
