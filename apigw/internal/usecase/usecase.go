package usecase

import (
	"github.com/nofendian17/gopkg/logger"
	airlineRepository "github.com/nofendian17/openota/apigw/internal/repository/airline"
	airportRepository "github.com/nofendian17/openota/apigw/internal/repository/airport"
	cityRepository "github.com/nofendian17/openota/apigw/internal/repository/city"
	countryRepository "github.com/nofendian17/openota/apigw/internal/repository/country"
	stateRepository "github.com/nofendian17/openota/apigw/internal/repository/state"
	"github.com/nofendian17/openota/apigw/internal/usecase/airline"
	"github.com/nofendian17/openota/apigw/internal/usecase/airport"
	"github.com/nofendian17/openota/apigw/internal/usecase/city"
	"github.com/nofendian17/openota/apigw/internal/usecase/country"
	"github.com/nofendian17/openota/apigw/internal/usecase/state"
	"time"

	"github.com/nofendian17/openota/apigw/internal/config"
	"github.com/nofendian17/openota/apigw/internal/infra/cache"
	"github.com/nofendian17/openota/apigw/internal/infra/database"
	"github.com/nofendian17/openota/apigw/internal/usecase/healthcheck"
)

type UseCase struct {
	Health  healthcheck.UseCase
	Country country.UseCase
	State   state.UseCase
	City    city.UseCase
	Airport airport.UseCase
	Airline airline.UseCase
}

// New creates a new instance of the UseCase struct, initializing it with the provided configuration and database.
func New(cfg *config.Config, logger logger.Logger, db *database.DB, cache cache.Client) *UseCase {
	healthUseCase := healthcheck.New(time.Now(), cfg, db, cache)

	countryRepo := countryRepository.New(db)
	countryUseCase := country.New(logger, countryRepo)

	stateRepo := stateRepository.New(db)
	stateUseCase := state.New(logger, stateRepo)

	cityRepo := cityRepository.New(db)
	cityUseCase := city.New(logger, cityRepo)

	airportRepo := airportRepository.New(db)
	airportUseCase := airport.New(logger, airportRepo)

	airlineRepo := airlineRepository.New(db)
	airlineUseCase := airline.New(logger, airlineRepo)

	return &UseCase{
		Health:  healthUseCase,
		Country: countryUseCase,
		State:   stateUseCase,
		City:    cityUseCase,
		Airport: airportUseCase,
		Airline: airlineUseCase,
	}
}
