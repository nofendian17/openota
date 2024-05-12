package healthcheck

import (
	"context"
	"github.com/nofendian17/openota/coresvc/internal/entity"
	"time"

	"github.com/nofendian17/openota/coresvc/internal/config"
	cacheClient "github.com/nofendian17/openota/coresvc/internal/infra/cache"
	"github.com/nofendian17/openota/coresvc/internal/infra/database"
)

type UseCase interface {
	Health(ctx context.Context) (*entity.HealthResponse, error)
	Readiness(ctx context.Context) (*entity.ReadinessResponse, error)
}

type useCase struct {
	startAt     time.Time
	config      *config.Config
	db          *database.DB
	cacheClient cacheClient.Client
}

func New(t time.Time, c *config.Config, db *database.DB, cacheClient cacheClient.Client) UseCase {
	return &useCase{
		startAt:     t,
		config:      c,
		db:          db,
		cacheClient: cacheClient,
	}
}
