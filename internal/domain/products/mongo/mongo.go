package mongo

import (
	"store-hexagon-architecture/internal/domain/products"
	"store-hexagon-architecture/internal/infrastructure"
	"store-hexagon-architecture/internal/utils/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo is contains functions for products db.
type Mongo struct {
	db       *mongo.Database
	products string
	otel     *infrastructure.Otel
	log      *logger.Logger
}

// RepoProvider to create new products db.
func RepoProvider(
	db *mongo.Database,
	products string,
	otel *infrastructure.Otel,
	log *logger.Logger) *Mongo {
	return &Mongo{
		db:       db,
		products: products,
		otel:     otel,
		log:      log,
	}
}

func (mg *Mongo) fromEntity(p products.ProductEntity) ProductMongo {
	return ProductMongo{
		Name:  p.Name,
		Stock: p.Stock,
	}
}

// Product is model database for product
type ProductMongo struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Stock int                `bson:"stock"`
}

func (p *ProductMongo) toEntity() *products.ProductEntity {
	return &products.ProductEntity{
		ID:    p.ID.Hex(),
		Name:  p.Name,
		Stock: p.Stock,
	}
}

func toEntities(p []ProductMongo) []*products.ProductEntity {
	productSlice := make([]*products.ProductEntity, len(p))
	for i, product := range p {
		productSlice[i] = product.toEntity()
	}
	return productSlice
}
