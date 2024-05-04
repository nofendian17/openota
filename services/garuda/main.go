package main

import (
	"context"
	"runtime"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/garuda/cmd"
	"github.com/nofendian17/openota/garuda/internal/config"
	"github.com/nofendian17/openota/garuda/internal/container"
	"github.com/nofendian17/openota/garuda/internal/delivery/rest"
	"github.com/nofendian17/openota/garuda/internal/infra/cache"
	"github.com/nofendian17/openota/garuda/internal/infra/database"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Initialize config
	cfg := config.New()

	// Initialize log
	l := logger.New(logger.Config{
		Output:  cfg.Logger.Output,
		Level:   cfg.Logger.Level,
		Service: cfg.Application.Name,
		Version: cfg.Application.Version,
	})

	// Initialize db
	db, err := database.New(cfg, l)
	if err != nil {
		panic(err)
	}

	// Initialize cache
	c, err := cache.New(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize container
	cntr := container.New(cfg, db, c, l)

	// Initialize rest server
	restServer := rest.New(cntr)
	srv := cmd.New(cntr, restServer)
	err = srv.StartRestServer(context.Background())
	if err != nil {
		panic(err)
	}
}
