package rest

func (s *server) routes() {
	// health routes
	s.router.HandleFunc("GET /health", s.handler.Health.Health())
	s.router.HandleFunc("GET /readiness", s.handler.Health.Readiness())
	s.router.HandleFunc("GET /ping", s.handler.Health.Ping())

	// country api routes
	s.router.HandleFunc("GET /api/v1/country", s.handler.Country.GetAll())
	s.router.HandleFunc("POST /api/v1/country", s.handler.Country.Create())
	s.router.HandleFunc("GET /api/v1/country/{countryID}", s.handler.Country.GetByID())
	s.router.HandleFunc("PUT /api/v1/country/{countryID}", s.handler.Country.Update())
	s.router.HandleFunc("DELETE /api/v1/country/{countryID}", s.handler.Country.Delete())
	s.router.HandleFunc("GET /api/v1/country/code/{countryCode}", s.handler.Country.GetByCode())

	// state api routes
	s.router.HandleFunc("GET /api/v1/state", s.handler.State.GetAll())
	s.router.HandleFunc("POST /api/v1/state", s.handler.State.Create())
	s.router.HandleFunc("GET /api/v1/state/{stateID}", s.handler.State.GetByID())
	s.router.HandleFunc("PUT /api/v1/state/{stateID}", s.handler.State.Update())
	s.router.HandleFunc("DELETE /api/v1/state/{stateID}", s.handler.State.Delete())
	s.router.HandleFunc("GET /api/v1/state/country/{countryID}", s.handler.State.GetByCountryID())

	// city api routes
	s.router.HandleFunc("GET /api/v1/city", s.handler.City.GetAll())
	s.router.HandleFunc("POST /api/v1/city", s.handler.City.Create())
	s.router.HandleFunc("GET /api/v1/city/{cityID}", s.handler.City.GetByID())
	s.router.HandleFunc("PUT /api/v1/city/{cityID}", s.handler.City.Update())
	s.router.HandleFunc("DELETE /api/v1/city/{cityID}", s.handler.City.Delete())
	s.router.HandleFunc("GET /api/v1/city/state/{stateID}", s.handler.City.GetByStateID())
}
