package products

import (
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func (hnd *productHandler) Create(c *fiber.Ctx) error {
	ctx, span := hnd.otel.Tracer().Start(c.UserContext(), "handler:product:create")
	defer span.End()

	request := new(model.CreateProductRequest)
	err := c.BodyParser(request)
	if err != nil {
		return responseutil.ResponseWithJSON(c, fiber.StatusBadRequest, nil, fiber.ErrBadRequest, nil)
	}

	result, code, err := hnd.productService.CreateProduct(ctx, *request)
	if err != nil {
		return responseutil.ResponseError(c, code, err)
	}

	return responseutil.ResponseWithJSON(c, code, result, err, nil)
}
