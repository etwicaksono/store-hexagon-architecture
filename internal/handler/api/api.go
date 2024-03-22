package api

import (
	"store-hexagon-architecture/internal/handler/api/products"
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(
	fiberApp *fiber.App,
	ph products.ProductHandler,
) {

	fiberApp.Get("/", func(c *fiber.Ctx) error { return nil })
	fiberApp.Get("/products", ph.List)
	fiberApp.Get("/product/:id", ph.Find)
	fiberApp.Post("/product", ph.Create)
	fiberApp.Put("/product/:id", ph.Update)
	fiberApp.Delete("/product/:id", ph.Delete)

	// Not found handler, this should be in bottom of routes
	fiberApp.Use(func(ctx *fiber.Ctx) error {
		return responseutil.ResponseError(ctx, fiber.ErrNotFound.Code, fiber.ErrNotFound)
	})
}
