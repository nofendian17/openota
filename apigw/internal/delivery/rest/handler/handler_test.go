package handler

import (
	"testing"

	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/container"
	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler/healthcheck"
	mockCacheClient "github.com/nofendian17/openota/apigw/internal/mocks/infra/cache"
	"github.com/nofendian17/openota/apigw/internal/usecase"
	"github.com/nofendian17/openota/apigw/pkg/logger"
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
	u := usecase.New(cfg, nil, c)

	cntr := &container.Container{
		Config:  cfg,
		UseCase: u,
		Logger:  l,
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
				Health: healthcheck.New(cntr),
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
