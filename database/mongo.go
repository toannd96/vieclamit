package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo struct
type Mongo struct {
	Db *mongo.Database
}

// CreateConn create connection to mongodb
func (m *Mongo) CreateConn() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	m.Db = client.Database("")
}
