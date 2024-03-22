package products

import (
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func (hnd *productHandler) Find(c *fiber.Ctx) error {
	ctx, span := hnd.otel.Tracer().Start(c.UserContext(), "handler:product:find")
	defer span.End()

	id := c.Params("id")

	result, code, err := hnd.productService.GetProduct(ctx, id)
	if err != nil {
		return responseutil.ResponseError(c, code, err)
	}

	return responseutil.ResponseWithJSON(c, code, result, err, nil)
}
