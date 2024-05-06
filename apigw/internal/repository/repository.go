package repository

import (
	cacheClient "github.com/nofendian17/openota/apigw/internal/infra/cache"
	"github.com/nofendian17/openota/apigw/internal/infra/database"
	cacheRepository "github.com/nofendian17/openota/apigw/internal/repository/cache"
	cityRepository "github.com/nofendian17/openota/apigw/internal/repository/city"
	countryRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
	stateRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
)

// Repository represents a repository entity.
type Repository struct {
	CacheRepository   cacheRepository.Repository // CacheRepository is the cache repository.
	CountryRepository countryRepository.Repository
	StateRepository   stateRepository.Repository
	CityRepository    cityRepository.Repository
}

// New creates a new instance of Repository.
func New(client cacheClient.Client, db *database.DB) *Repository {
	cacheRepo := cacheRepository.New(client)
	countryRepo := countryRepository.New(db)
	stateRepo := stateRepository.New(db)
	cityRepo := cityRepository.New(db)

	return &Repository{
		CacheRepository:   cacheRepo,
		CountryRepository: countryRepo,
		StateRepository:   stateRepo,
		CityRepository:    cityRepo,
	}
}
