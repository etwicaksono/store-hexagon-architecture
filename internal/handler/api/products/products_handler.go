package products

import (
	"store-hexagon-architecture/internal/infrastructure"
	"store-hexagon-architecture/internal/service/products"
	"store-hexagon-architecture/internal/utils/logger"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	List(c *fiber.Ctx) error
	Find(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type productHandler struct {
	log            *logger.Logger
	otel           infrastructure.Otel
	productService products.Service
}

func HandlerProvider(
	log *logger.Logger,
	otel infrastructure.Otel,
	productSvc products.Service,
) ProductHandler {
	return &productHandler{
		log:            log,
		otel:           otel,
		productService: productSvc,
	}
}
