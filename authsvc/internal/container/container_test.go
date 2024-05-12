package container

import (
	"github.com/nofendian17/gopkg/validator"
	"reflect"
	"testing"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/authsvc/internal/config"
	"github.com/nofendian17/openota/authsvc/internal/infra/cache"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
	"github.com/nofendian17/openota/authsvc/internal/usecase"
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

	v := validator.NewValidator()
	type args struct {
		cfg *config.Config
		db  *database.DB
		c   cache.Client
		l   logger.Logger
		v   *validator.CustomValidator
	}
	tests := []struct {
		name string
		args args
		want *Container
	}{
		{
			name: "success",
			args: args{
				cfg: cfg,
				db:  nil,
				c:   nil,
				l:   l,
				v:   v,
			},
			want: &Container{
				Config:  cfg,
				UseCase: usecase.New(cfg, l, nil, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.db, tt.args.c, tt.args.l, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				assert.IsType(t, tt.want, got)
			}
		})
	}
}
