package products

import (
	"context"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/infrastructure"
	"store-hexagon-architecture/internal/model"

	"github.com/go-playground/validator/v10"
)

// Service contains functions for service.
type Service interface {
	GetProduct(ctx context.Context, objectID string) (*model.Product, int, error)
	GetProductList(ctx context.Context, req model.GetProductListRequest) ([]*model.Product, *model.Pagination, int, error)
	CreateProduct(ctx context.Context, data model.CreateProductRequest) (*model.Product, int, error)
	UpdateProduct(ctx context.Context, objectID string, updateData model.UpdateProductRequest) (*model.Product, int, error)
	DeleteProduct(ctx context.Context, objectID string) (int, error)
}

type service struct {
	otel     *infrastructure.Otel
	vld      *validator.Validate
	products products.Repository
}

// New to create new service.
func ServiceProvider(
	otel *infrastructure.Otel,
	vld *validator.Validate,
	products products.Repository,
) Service {
	return &service{
		otel:     otel,
		vld:      vld,
		products: products,
	}
}
