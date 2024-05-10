package handler

import (
	"github.com/nofendian17/gopkg/validator"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/airport"
	"testing"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/city"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/country"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/healthcheck"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/state"
	mockCacheClient "github.com/nofendian17/openota/apigw/internal/mocks/infra/cache"
	"github.com/nofendian17/openota/apigw/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	l := logger.New(logger.Config{
		Output:  cfg.Logger.Output,
		Level:   cfg.Logger.Level,
		Service: cfg.Application.Name,
		Version: cfg.Application.Version,
	})

	c := &mockCacheClient.Client{}
	u := usecase.New(cfg, l, nil, c)
	v := validator.NewValidator()

	cntr := &container.Container{
		Config:    cfg,
		UseCase:   u,
		Logger:    l,
		Validator: v,
	}

	type args struct {
		c *container.Container
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				c: cntr,
			},
			want: &Handler{
				Health:  healthcheck.New(cntr),
				Country: country.New(cntr),
				State:   state.New(cntr),
				City:    city.New(cntr),
				Airport: airport.New(cntr),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.c)
			assert.Equal(t, tt.want, got)
		})
	}
}
