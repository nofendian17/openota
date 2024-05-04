package healthcheck

import (
	"testing"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/lion/internal/config"
	"github.com/nofendian17/openota/lion/internal/container"
	mockCacheClient "github.com/nofendian17/openota/lion/internal/mocks/infra/cache"
	"github.com/nofendian17/openota/lion/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
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
		want Handler
	}{
		{
			name: "should return a new handler",
			args: args{
				c: cntr,
			},
			want: &handler{
				config:  cfg,
				useCase: u.Health,
				logger:  l,
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
