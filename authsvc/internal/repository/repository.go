package repository

import (
	cacheClient "github.com/nofendian17/openota/authsvc/internal/infra/cache"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
	cacheRepository "github.com/nofendian17/openota/authsvc/internal/repository/cache"
)

// Repository represents a repository entity.
type Repository struct {
	CacheRepository cacheRepository.Repository // CacheRepository is the cache repository.
}

// New creates a new instance of Repository.
func New(client cacheClient.Client, db *database.DB) *Repository {
	cacheRepo := cacheRepository.New(client)

	return &Repository{
		CacheRepository: cacheRepo,
	}
}
