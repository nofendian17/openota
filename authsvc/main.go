package main

import (
	"context"
	"github.com/nofendian17/openota/authsvc/internal/delivery/grpc"
	"runtime"

	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/gopkg/validator"
	"github.com/nofendian17/openota/authsvc/cmd"
	"github.com/nofendian17/openota/authsvc/internal/config"
	"github.com/nofendian17/openota/authsvc/internal/container"
	"github.com/nofendian17/openota/authsvc/internal/infra/cache"
	"github.com/nofendian17/openota/authsvc/internal/infra/database"
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

	// initialize validator
	v := validator.NewValidator()

	// Initialize db
	db, err := database.New(cfg, l)
	if err != nil {
		panic(err)
	}

	// Do migration
	err = db.GormDB.AutoMigrate()
	if err != nil {
		panic(err)
	}

	// Initialize cache
	c, err := cache.New(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize container
	cntr := container.New(cfg, db, c, l, v)

	// Initialize grpc server
	grpcServer := grpc.New(cntr)
	srv := cmd.New(cntr, grpcServer)
	err = srv.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
