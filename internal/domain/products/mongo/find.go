package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *DB) GetProductByID(ctx context.Context, objectID string) (*products.Product, int, error) {
	startTime := time.Now()

	ctx, span := d.otel.Tracer().Start(ctx, "db:GetProductByID")
	defer span.End()

	id, _ := primitive.ObjectIDFromHex(objectID)
	filter := bson.M{
		"id":         id,
		"deleted_at": bson.M{"$exists": false},
	}

	var pr Product
	err := d.db.Collection(d.products).FindOne(ctx, filter).Decode(&pr)
	if err == mongo.ErrNoDocuments {
		return nil, http.StatusNotFound, errors.ErrNotFoundProduct
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	d.log.Info(fmt.Sprintf("Execution Time (Get Product By Id): %s\n", executionTime))
	return pr.toEntity(), http.StatusOK, nil
}
