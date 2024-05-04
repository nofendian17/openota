package usecase

import (
	"time"

	"github.com/nofendian17/openota/sabre/internal/config"
	"github.com/nofendian17/openota/sabre/internal/infra/cache"
	"github.com/nofendian17/openota/sabre/internal/infra/database"
	"github.com/nofendian17/openota/sabre/internal/usecase/healthcheck"
)

type UseCase struct {
	Health healthcheck.UseCase
}

// New creates a new instance of the UseCase struct, initializing it with the provided configuration and database.
func New(cfg *config.Config, db *database.DB, cache cache.Client) *UseCase {
	healthUseCase := healthcheck.New(time.Now(), cfg, db, cache)
	return &UseCase{
		Health: healthUseCase,
	}
}
