package products

import (
	"context"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/model"
)

func (s *service) CreateProduct(ctx context.Context, data model.CreateProductRequest) (*model.Product, int, error) {
	_, span := s.otel.Tracer().Start(ctx, "service:product:create")
	defer span.End()

	if err := s.vld.Struct(&data); err != nil {
		return nil, http.StatusBadRequest, err
	}

	product, code, err := s.products.Create(ctx, products.ProductEntity{
		Name:  data.Name,
		Stock: data.Stock,
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
