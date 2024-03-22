package products

import (
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func (hnd *productHandler) Update(c *fiber.Ctx) error {
	ctx, span := hnd.otel.Tracer().Start(c.UserContext(), "handler:product:update")
	defer span.End()

	id := c.Params("id")
	request := new(model.UpdateProductRequest)
	err := c.BodyParser(request)
	if err != nil {
		return responseutil.ResponseWithJSON(c, fiber.StatusBadRequest, nil, fiber.ErrBadRequest, nil)
	}

	result, code, err := hnd.productService.UpdateProduct(ctx, id, *request)
	if err != nil {
		return responseutil.ResponseError(c, code, err)
	}

	return responseutil.ResponseWithJSON(c, code, result, err, nil)
}
