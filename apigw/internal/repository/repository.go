package repository

import (
	cacheClient "github.com/nofendian17/openota/apigw/internal/infra/cache"
	cacheRepository "github.com/nofendian17/openota/apigw/internal/repository/cache"
)

// Repository represents a repository entity.
type Repository struct {
	CacheRepository cacheRepository.Repository // CacheRepository is the cache repository.
}

// New creates a new instance of Repository.
func New(client cacheClient.Client) *Repository {
	cacheRepo := cacheRepository.New(client)
	return &Repository{
		CacheRepository: cacheRepo,
	}
}
