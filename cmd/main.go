package main

import (
	"context"
	"os"
	"os/signal"
	"store-hexagon-architecture/config"
	"store-hexagon-architecture/internal/connection"
	productMongo "store-hexagon-architecture/internal/domain/products/mongo"
	"store-hexagon-architecture/internal/handler/api"
	productHnd "store-hexagon-architecture/internal/handler/api/products"
	"store-hexagon-architecture/internal/infrastructure"
	"store-hexagon-architecture/internal/pkg/http"
	productSvc "store-hexagon-architecture/internal/service/products"
	"store-hexagon-architecture/internal/utils/logger"
	"store-hexagon-architecture/internal/utils/validatorutil"
	"syscall"
	"time"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	mongoDB, err := conn.NewMongoDB()
	if err != nil {
		return
	}

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

	// Initialize web server
	httpServer := http.New(http.Config{
		Host:         cfg.GetString("app.host"),
		Port:         cfg.GetString("app.port"),
		IdleTimeout:  time.Duration(cfg.GetInt("fiber.idleTimeout")) * time.Second,
		WriteTimeout: time.Duration(cfg.GetInt("fiber.writeTimeout")) * time.Second,
		ReadTimeout:  time.Duration(cfg.GetInt("fiber.readTimeout")) * time.Second,
		Prefork:      cfg.GetBool("fiber.prefork"),
	})
	router := httpServer.Router()
	router.Use(recover.New(recover.Config{
		EnableStackTrace: cfg.GetBool("fiber.enableStackTrace"),
	}))
	router.Use(otelfiber.Middleware())

	// Initialize  validator
	vld := validatorutil.GetValidator()

	// Initialize repository
	productRepo := productMongo.RepoProvider(mongoDB, "products", &otl, lgr)

	// Initialize service
	productService := productSvc.ServiceProvider(&otl, vld, productRepo)

	// Initialize handler
	productHandler := productHnd.HandlerProvider(lgr, otl, productService)

	// Register api route.
	api.SetRoutes(router, productHandler)

	// Run web server.
	httpServerChan := httpServer.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-httpServerChan:
		if err != nil {
			return
		}
	case <-sigChan:
	}
}
