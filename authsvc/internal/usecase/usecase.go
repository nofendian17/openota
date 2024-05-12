package usecase

import (
	"github.com/nofendian17/gopkg/logger"
	"time"

	"github.com/nofendian17/openota/authsvc/internal/config"
	"github.com/nofendian17/openota/authsvc/internal/infra/cache"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
	"github.com/nofendian17/openota/authsvc/internal/usecase/healthcheck"
)

type UseCase struct {
	Health healthcheck.UseCase
}

// New creates a new instance of the UseCase struct, initializing it with the provided configuration and database.
func New(cfg *config.Config, logger logger.Logger, db *database.DB, cache cache.Client) *UseCase {
	healthUseCase := healthcheck.New(time.Now(), cfg, db, cache)

	return &UseCase{
		Health: healthUseCase,
	}
}
