package repository

import (
	"github.com/nofendian17/openota/apigw/internal/infra/database"
	mockDB "github.com/nofendian17/openota/apigw/internal/mocks/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/infra/cache"
	cacheRepository "github.com/nofendian17/openota/apigw/internal/repository/cache"
	cityRepository "github.com/nofendian17/openota/apigw/internal/repository/city"
	countryRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
	stateRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	cacheClient, _ := cache.New(cfg)
	_mockDB, _, _ := mockDB.New(cfg.Database.Driver)

	type args struct {
		client cache.Client
		db     *database.DB
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			name: "success",
			args: args{
				client: cacheClient,
				db:     (*database.DB)(_mockDB),
			},
			want: &Repository{
				CacheRepository:   cacheRepository.New(cacheClient),
				CountryRepository: countryRepository.New((*database.DB)(_mockDB)),
				StateRepository:   stateRepository.New((*database.DB)(_mockDB)),
				CityRepository:    cityRepository.New((*database.DB)(_mockDB)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.client, tt.args.db)
			assert.Equal(t, tt.want, got)
		})
	}
}
