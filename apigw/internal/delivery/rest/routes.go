package rest

func (s *server) routes() {
	// health routes
	s.router.HandleFunc("GET /health", s.handler.Health.Health())
	s.router.HandleFunc("GET /readiness", s.handler.Health.Readiness())
	s.router.HandleFunc("GET /ping", s.handler.Health.Ping())

	s.router.HandleFunc("GET /api/v1/country/{countryID}", s.handler.Country.GetByID())
	s.router.HandleFunc("GET /api/v1/country/code/{countryCode}", s.handler.Country.GetByCode())
}
