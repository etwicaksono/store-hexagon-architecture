package products

import (
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils/responseutil"

	"github.com/gofiber/fiber/v2"
)

func (hnd *productHandler) List(c *fiber.Ctx) error {
	ctx, span := hnd.otel.Tracer().Start(c.UserContext(), "handler:product:list")
	defer span.End()

	request := new(model.GetProductListRequest)
	err := c.QueryParser(request)
	if err != nil {
		return responseutil.ResponseWithJSON(c, fiber.StatusBadRequest, nil, fiber.ErrBadRequest, nil)
	}

	if request.Page <= 0 {
		request.Page = 1
	}

	if request.Limit <= 0 {
		request.Limit = 5
	}

	result, pagination, code, err := hnd.productService.GetProductList(ctx, *request)
	if err != nil {
		return responseutil.ResponseError(c, code, err)
	}

	return responseutil.ResponseWithJSON(c, code, result, err, nil, pagination)
}
