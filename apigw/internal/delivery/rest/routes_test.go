package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nofendian17/openota/apigw/internal/delivery/rest/handler"
	mockAirlineHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/airline"
	mockAirportHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/airport"
	mockCityHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/city"
	mockCountryHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/country"
	mockHealthCheckHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/healthcheck"
	mockStateHandler "github.com/nofendian17/openota/apigw/internal/mocks/delivery/rest/handler/state"

	"github.com/stretchr/testify/assert"
)

func TestServer_routes(t *testing.T) {
	_mockHealthHandler := &mockHealthCheckHandler.Handler{}
	_mockHealthHandler.On("Health").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	}).Once()
	_mockHealthHandler.On("Readiness").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockHealthHandler.On("Ping").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})

	_mockCountryHandler := &mockCountryHandler.Handler{}
	_mockCountryHandler.On("GetByID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCountryHandler.On("GetByCode").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCountryHandler.On("GetAll").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})

	_mockCountryHandler.On("Create").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}
	})
	_mockCountryHandler.On("Update").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCountryHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler := &mockStateHandler.Handler{}
	_mockStateHandler.On("GetByID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler.On("GetAll").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler.On("GetByCountryID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler.On("Update").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockStateHandler.On("Create").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}
	})
	_mockStateHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})

	_mockCityHandler := &mockCityHandler.Handler{}
	_mockCityHandler.On("Create").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}
	})
	_mockCityHandler.On("Update").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCityHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCityHandler.On("GetAll").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCityHandler.On("GetByID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockCityHandler.On("GetByStateID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirportHandler := &mockAirportHandler.Handler{}
	_mockAirportHandler.On("Create").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}
	})
	_mockAirportHandler.On("Update").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirportHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirportHandler.On("GetAll").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirportHandler.On("GetByID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirportHandler.On("GetByCode").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirlineHandler := &mockAirlineHandler.Handler{}
	_mockAirlineHandler.On("GetAll").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirlineHandler.On("Create").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}
	})
	_mockAirlineHandler.On("GetByID").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirlineHandler.On("Update").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirlineHandler.On("Delete").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})
	_mockAirlineHandler.On("GetByCode").Return(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}
	})

	// Create a new instance of handler.Handler
	h := &handler.Handler{
		Health:  _mockHealthHandler,
		Country: _mockCountryHandler,
		State:   _mockStateHandler,
		City:    _mockCityHandler,
		Airport: _mockAirportHandler,
		Airline: _mockAirlineHandler,
	}

	// You may need to replace this with a mock or real implementation
	// Create a new instance of Server with the mock handler
	s := &server{
		router:  http.NewServeMux(),
		handler: h,
	}

	// Call the routes method to register routes
	s.routes()

	// Create a test server
	testServer := httptest.NewServer(s.router)
	defer testServer.Close()

	// Make GET requests to the registered routes and check if the status codes are OK
	healthResp, err := http.Get(testServer.URL + "/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, healthResp.StatusCode)

	readinessResp, err := http.Get(testServer.URL + "/readiness")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, readinessResp.StatusCode)

	pingResp, err := http.Get(testServer.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, pingResp.StatusCode)
}
