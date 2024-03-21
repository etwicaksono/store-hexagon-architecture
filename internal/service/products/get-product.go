package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

func (s *service) GetProduct(ctx context.Context, objectID string) (*model.Product, int, error) {
	_, span := s.otel.Tracer().Start(ctx, "service:GetProduct")
	defer span.End()

	product, code, err := s.products.Find(ctx, objectID)
	if err != nil {

		return nil, code, err
	}

	return &model.Product{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
	}, code, nil
}
