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

// New to create new products db.
func New(
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

func (db *Mongo) fromEntity(p products.Product) Product {
	return Product{
		Name:  p.Name,
		Stock: p.Stock,
	}
}

// Product is model database for product
type Product struct {
	ID    primitive.ObjectID `bson:"id,omitempty"`
	Name  string             `bson:"name"`
	Stock int                `bson:"stock"`
}

func (p *Product) toEntity() *products.Product {
	return &products.Product{
		ID:    p.ID.Hex(),
		Name:  p.Name,
		Stock: p.Stock,
	}
}

func toEntities(p []Product) []*products.Product {
	productSlice := make([]*products.Product, len(p))
	for i, product := range p {
		productSlice[i] = product.toEntity()
	}
	return productSlice
}