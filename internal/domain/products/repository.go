package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

type Repository interface {
	Find(ctx context.Context, objectID string) (data *Product, statusCode int, err error)
	List(ctx context.Context, req model.GetProductRequest) (data []*Product, pgn *model.Pagination, statusCode int, err error)
	Create(ctx context.Context, product Product) (data *Product, statusCode int, err error)
	Update(ctx context.Context, id string, updateData model.UpdateProductRequest) (data *Product, statusCode int, err error)
	Delete(ctx context.Context, id string) (statusCode int, err error)
}
