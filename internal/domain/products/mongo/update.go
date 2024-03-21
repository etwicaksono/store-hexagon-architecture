package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils/errorutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mg *Mongo) Update(ctx context.Context, objectID string, updateData model.UpdateProductRequest) (data *products.ProductEntity, statusCode int, err error) {
	startTime := time.Now()
	ctx, span := mg.otel.Tracer().Start(ctx, "db:product:update")
	defer span.End()

	id, err := primitive.ObjectIDFromHex(objectID)
	if err != nil {
		return nil, http.StatusBadRequest, errorutil.ErrNotFoundProduct
	}

	filter := bson.M{
		"id": id,
	}

	update := bson.M{}
	if updateData.Name != "" {
		update["name"] = updateData.Name
	}
	if updateData.Stock != nil {
		update["stock"] = *updateData.Stock
	}

	updateQuery := bson.M{"$set": update}

	var pr ProductMongo
	err = mg.db.Collection(mg.products).FindOneAndUpdate(ctx, filter, updateQuery, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&pr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, errorutil.ErrNotFoundProduct
		}
		return nil, http.StatusInternalServerError, errorutil.ErrInternalDB
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	mg.log.Info(fmt.Sprintf("Execution Time (db:product:update): %s\n", executionTime))

	return pr.toEntity(), http.StatusOK, nil
}
