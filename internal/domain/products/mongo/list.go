package mongo

import (
	"context"
	"fmt"
	"net/http"
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/model"
	"store-hexagon-architecture/internal/utils/errorhelper"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mg *Mongo) List(ctx context.Context, req model.GetProductsRequest) (data []*products.ProductEntity, pgn *model.Pagination, statusCode int, err error) {
	startTime := time.Now()

	ctx, span := mg.otel.Tracer().Start(ctx, "db:product:find")
	defer span.End()

	filter := make(map[string]interface{})

	if trimmedName := strings.TrimSpace(req.Name); trimmedName != "" {
		filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: trimmedName, Options: "i"}}
	}

	limit := int64(req.Limit)
	skip := int64(req.Page*req.Limit - req.Limit)
	options := options.Find()

	options.SetLimit(limit)
	options.Skip = &skip

	cur, err := mg.db.Collection(mg.products).Find(ctx, filter, options)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, errorhelper.ErrInternalDB
	}
	defer cur.Close(ctx)

	var products []ProductMongo
	for cur.Next(ctx) {
		var product ProductMongo
		err := cur.Decode(&product)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, errorhelper.ErrInternalDB
		}
		products = append(products, product)
	}
	if err := cur.Err(); err != nil {
		return nil, nil, http.StatusInternalServerError, errorhelper.ErrInternalDB
	}

	total, err := mg.db.Collection(mg.products).CountDocuments(ctx, filter)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, errorhelper.ErrInternalDB
	}
	defer cur.Close(ctx)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	mg.log.Info(fmt.Sprintf(" Execution Time (db:product:find): %s\n", executionTime))

	return toEntities(products), &model.Pagination{
		Total:       int(total),
		Limit:       req.Limit,
		CurrentPage: req.Page,
		LastPage:    0,
	}, http.StatusOK, nil
}
