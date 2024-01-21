package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoInstance *mongo.Database

func InitMongoDb() {
	uri := "mongodb://localhost/?retryWrites=true&w=majority"
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	MongoInstance = client.Database("ygo_deck_master")
}
