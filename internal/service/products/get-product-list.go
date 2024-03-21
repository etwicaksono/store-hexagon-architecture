package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

func (s *service) GetProductList(ctx context.Context, req model.GetProductListRequest) ([]*model.Product, *model.Pagination, int, error) {
	_, span := s.otel.Tracer().Start(ctx, "service:product:getlist")
	defer span.End()

	products, pagination, code, err := s.products.List(ctx, req)
	if err != nil {
		return nil, nil, code, err
	}

	prds := make([]*model.Product, len(products))
	for i, product := range products {
		productDTO := &model.Product{
			ID:    product.ID,
			Name:  product.Name,
			Stock: product.Stock,
		}
		prds[i] = productDTO
	}

	return prds, pagination, code, nil
}
