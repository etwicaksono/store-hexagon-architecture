package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/utils/errorutil"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mg *Mongo) Create(ctx context.Context, product products.ProductEntity) (data *products.ProductEntity, statusCode int, err error) {
	startTime := time.Now()

	ctx, span := mg.otel.Tracer().Start(ctx, "db:product:create")
	defer span.End()

	prd := mg.fromEntity(product)
	res, err := mg.db.Collection(mg.products).InsertOne(ctx, prd)
	if err != nil {
		mg.log.Error(err.Error())
		return nil, http.StatusInternalServerError, errorutil.ErrInternalDB
	}

	prd.ID = res.InsertedID.(primitive.ObjectID)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	mg.log.Info(fmt.Sprintf("Execution Time (db:product:create): %s\n", executionTime))
	return prd.toEntity(), http.StatusOK, nil
}
