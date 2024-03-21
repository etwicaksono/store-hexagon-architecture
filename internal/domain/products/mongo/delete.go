package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/utils/errorutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mg *Mongo) Delete(ctx context.Context, objectID string) (statusCode int, err error) {
	startTime := time.Now()

	ctx, span := mg.otel.Tracer().Start(ctx, "db:product:delete")
	defer span.End()

	id, err := primitive.ObjectIDFromHex(objectID)
	if err != nil {
		return http.StatusBadRequest, errorutil.ErrNotFoundProduct
	}

	filter := bson.M{"_id": id}

	// Ganti UpdateOne dengan DeleteOne
	result, err := mg.db.Collection(mg.products).DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return http.StatusNotFound, errorutil.ErrNotFoundProduct
		}
		return http.StatusInternalServerError, errorutil.ErrInternalDB
	}

	if result.DeletedCount == 0 {
		return http.StatusNotFound, errorutil.ErrNotFoundProduct
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	mg.log.Info(fmt.Sprintf("Execution Time (db:product:delete): %s\n", executionTime))
	return http.StatusOK, nil
}
