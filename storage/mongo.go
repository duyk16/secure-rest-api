package storage

import (
	"context"
	"log"
	"secure-rest-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func Init() {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(config.ServerConfig.Storage.Uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Println("Connect to MongoDB fail")
		return
	}

	Database = client.Database(config.ServerConfig.Storage.Name)
	log.Println("Connected to MongoDB")
}
