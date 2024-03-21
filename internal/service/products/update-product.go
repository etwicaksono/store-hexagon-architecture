package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

func (s *service) UpdateProduct(ctx context.Context, objectID string, updateData model.UpdateProductRequest) (*model.Product, int, error) {
	_, span := s.otel.Tracer().Start(ctx, "service:UpdateProduct")
	defer span.End()

	product, code, err := s.products.Update(ctx, objectID, model.UpdateProductRequest{
		Name:  updateData.Name,
		Stock: updateData.Stock,
	})
	if err != nil {
		return nil, code, err
	}

	return &model.Product{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
	}, code, nil
}
