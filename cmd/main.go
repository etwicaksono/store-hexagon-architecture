package main

import (
	"context"
	"store-hexagon-architecture/config"
	"store-hexagon-architecture/internal/connection"
	"store-hexagon-architecture/internal/infrastructure"
	"store-hexagon-architecture/internal/utils/logger"
)

func main() {
	// Get config.
	cfg := config.GetConfig()

	// Get logger
	lgr := logger.NewLogger(cfg)

	// Init db
	conn := connection.Connection{
		Config: cfg,
		Logger: lgr,
	}
	_, err := conn.NewMongoDB()
	if err != nil {
		return
	}

	// Init repository
	// Init service

	ctx := context.Background()

	// Init Opentelemetry.
	otl := infrastructure.Otel{
		Config: cfg,
		Logger: lgr,
	}
	otelShutdown, err := otl.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = otelShutdown(ctx)
	}()
}
