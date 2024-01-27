package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoInstance *mongo.Database

func InitMongoDb() {
	uri := "mongodb://192.168.48.1/?retryWrites=true&w=majority"
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	MongoInstance = client.Database("ygo_deck_master")
}
