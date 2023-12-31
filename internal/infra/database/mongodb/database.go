package mongodb

import (
	"blocklist/internal/config"
	"blocklist/internal/infra/database"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instanace *mongo.Client
var db *mongo.Database

// GetInstance return copy of db session
func GetInstance(c *config.DatabaseConfig) *mongo.Client {

	if instanace == nil {
		var err error
		instanace, err = Connect(c)
		db = instanace.Database(c.Dbname)
		if err != nil {
			panic(err)
		}
	}
	return instanace
}

func GetCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}


// Connect to the database
func Connect(config *config.DatabaseConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(database.GenerateMongoConnectionString(config))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
