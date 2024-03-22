package products

import (
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func (hnd *productHandler) Delete(c *fiber.Ctx) error {
	ctx, span := hnd.otel.Tracer().Start(c.UserContext(), "handler:product:delete")
	defer span.End()

	id := c.Params("id")

	code, err := hnd.productService.DeleteProduct(ctx, id)
	if err != nil {
		return responseutil.ResponseError(c, code, err)
	}

	return responseutil.ResponseWithJSON(c, code, fiber.Map{"message": "Product deleted successfully"}, err, nil)
}
