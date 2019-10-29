package storage

import (
	"context"
	"log"

	"github.com/duyk16/secure-rest-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var User *mongo.Collection

func Init() {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(config.ServerConfig.Storage.Uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Println("Connect to MongoDB fail")
		return
	}
	log.Println("Connected to MongoDB")

	Database = client.Database(config.ServerConfig.Storage.Name)
	User = Database.Collection("users")
	createIndexes()
}

func createIndexes() {
	User.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})
}
