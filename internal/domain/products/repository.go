package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

type Repository interface {
	GetProductByID(ctx context.Context, id string) (*Product, int, error)
	GetCompanies(ctx context.Context, data model.GetProductRequest) ([]*Product, *model.Pagination, int, error)
	CreateProduct(ctx context.Context, product Product) (*Product, int, error)
	UpdateProduct(ctx context.Context, id string, updateData model.UpdateProductRequest) (*Product, int, error)
	DeleteProduct(ctx context.Context, id string) (int, error)
}
