package connection

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func (c *Connection) NewMongoDB() (*mongo.Database, error) {

	// setup mongodb uri
	var dns string
	environment := c.Config.GetString("environment")
	dbAddress := c.Config.GetString("db.address")
	dbName := c.Config.GetString("db.name")
	dbMaxConnectionOpen := uint64(c.Config.GetInt("db.maxConnectionOpen"))
	dbMaxConnectionIdle := time.Duration(c.Config.GetInt("db.maxConnectionIdle")) * time.Second

	switch environment {
	case "local", "development":
		dns = fmt.Sprintf("mongodb://%s", dbAddress)
	default:
	}

	dbOptions := options.Client()
	dbOptions.Monitor = otelmongo.NewMonitor()
	dbOptions.MaxConnecting = &dbMaxConnectionOpen
	dbOptions.MaxConnIdleTime = &dbMaxConnectionIdle
	dbOptions.ApplyURI(dns)

	// Prepare dns and open connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connection to mongo database
	dbClient, err := mongo.Connect(ctx, dbOptions)
	if err != nil {
		c.Logger.Error("Failed to connect to database: ", err.Error())
		return nil, errors.Join(err)
	}

	c.Logger.Info("success to connect to database ")

	// check ping database
	err = dbClient.Ping(context.Background(), nil)
	if err != nil {
		c.Logger.Error("Failed to ping database: ", err.Error())
		return nil, err
	}

	c.Logger.Info("success ping to connect to database ")

	return dbClient.Database(dbName), nil

}
