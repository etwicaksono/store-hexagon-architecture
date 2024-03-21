package products

import (
	"context"
	"net/http"
)

func (s *service) DeleteProduct(ctx context.Context, objectID string) (int, error) {
	_, span := s.otel.Tracer().Start(ctx, "service:product:delete")
	defer span.End()

	code, err := s.products.Delete(ctx, objectID)
	if err != nil {
		return code, err
	}

	return http.StatusOK, nil
}
