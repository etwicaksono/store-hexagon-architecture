package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/utils/errorhelper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mg *Mongo) Find(ctx context.Context, objectID string) (data *products.ProductEntity, statusCode int, err error) {
	startTime := time.Now()

	ctx, span := mg.otel.Tracer().Start(ctx, "db:product:find")
	defer span.End()

	id, _ := primitive.ObjectIDFromHex(objectID)
	filter := bson.M{
		"id":         id,
		"deleted_at": bson.M{"$exists": false},
	}

	var pr ProductMongo
	err = mg.db.Collection(mg.products).FindOne(ctx, filter).Decode(&pr)
	if err == mongo.ErrNoDocuments {
		return nil, http.StatusNotFound, errorhelper.ErrNotFoundProduct
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	mg.log.Info(fmt.Sprintf("Execution Time (db:product:find): %s\n", executionTime))
	return pr.toEntity(), http.StatusOK, nil
}
