package products

import (
	"context"
	"store-hexagon-architecture/internal/model"
)

type Repository interface {
	Find(ctx context.Context, objectID string) (data *ProductEntity, statusCode int, err error)
	List(ctx context.Context, req model.GetProductListRequest) (data []*ProductEntity, pgn *model.Pagination, statusCode int, err error)
	Create(ctx context.Context, product ProductEntity) (data *ProductEntity, statusCode int, err error)
	Update(ctx context.Context, objectID string, updateData model.UpdateProductRequest) (data *ProductEntity, statusCode int, err error)
	Delete(ctx context.Context, objectID string) (statusCode int, err error)
}
