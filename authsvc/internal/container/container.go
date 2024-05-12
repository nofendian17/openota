package container

import (
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/gopkg/validator"
	"github.com/nofendian17/openota/authsvc/internal/config"
	"github.com/nofendian17/openota/authsvc/internal/infra/cache"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
	"github.com/nofendian17/openota/authsvc/internal/usecase"
)

type Container struct {
	Config    *config.Config
	UseCase   *usecase.UseCase
	Logger    logger.Logger
	Validator *validator.CustomValidator
}

// New initializes and returns a new Container with the given configuration.
func New(cfg *config.Config, db *database.DB, c cache.Client, l logger.Logger, v *validator.CustomValidator) *Container {
	return &Container{
		Config:    cfg,
		UseCase:   usecase.New(cfg, l, db, c),
		Logger:    l,
		Validator: v,
	}
}
