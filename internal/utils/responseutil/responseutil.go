package responseutil

import (
	"net/http"
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils"
	"store-hexagon-architecture/internal/utils/validatorutil"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ResponseWithJSON to write response with JSON format.
func ResponseWithJSON(c *fiber.Ctx, code int, data interface{}, err error, validationErr map[string]any, pagination ...*model.Pagination) error {
	response := model.Response{
		Status:  code,
		Message: strings.ToLower(http.StatusText(code)),
	}
	if len(pagination) > 0 && pagination[0] != nil {
		response.Pagination = pagination[0]
		if response.Pagination.CurrentPage <= 0 {
			response.Pagination.CurrentPage = 1
		}
		response.Pagination.LastPage = utils.RoundUp(float64(response.Pagination.Total) / float64(response.Pagination.Limit))
		if response.Pagination.LastPage <= 0 {
			response.Pagination.LastPage = 1
		}
	}

	response.Data = data
	if err != nil {
		response.Message = err.Error()
	}

	if len(validationErr) > 0 {
		response.Errors = validationErr
	}

	return c.Status(code).JSON(response)
}

func ResponseError(c *fiber.Ctx, code int, err error) error {
	switch code {
	case fiber.StatusBadRequest:
		{
			errValidation := validatorutil.GenerateErrorMessage(err)
			return ResponseWithJSON(c, code, nil, fiber.ErrBadRequest, errValidation)
		}
	default:
		{
			return ResponseWithJSON(c, code, nil, err, nil)
		}
	}
}
