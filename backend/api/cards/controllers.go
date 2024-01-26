package cards

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	Database "ygocarddb/database"
	models "ygocarddb/models"
	Http "ygocarddb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	findOptions := Http.Pagination(r)

	search := r.URL.Query().Get("search")
	var cursor *mongo.Cursor
	var err error

	if search != "" {
		findOptions.SetSort(bson.D{{Key: "name", Value: 1}})

		filter := bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "fr.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "de.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "it.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "pt.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "es.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
			}},
		}

		cursor, err = coll.Find(context.TODO(), filter, findOptions)
	} else {
		cursor, err = coll.Find(context.TODO(), bson.D{{}}, findOptions)
	}

	if err != nil {
		log.Panic(err)
	}

	var cards []models.Card

	if err = cursor.All(context.TODO(), &cards); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(cards)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	id, err := Http.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}

	var card models.Card
	cursor := coll.FindOne(context.TODO(), bson.D{
		{Key: "id", Value: id},
	})

	if cursor.Err() != nil {
		log.Panic(cursor.Err())
	}

	cursor.Decode(&card)

	json.NewEncoder(w).Encode(card)
}

func GetCardImage(w http.ResponseWriter, r *http.Request) {
	id, err := Http.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}
	filePath := "./assets/cards_small/" + strconv.Itoa(id) + ".jpg"
	Http.SendImage(w, filePath)
}

func GetCardImageBig(w http.ResponseWriter, r *http.Request) {
	id, err := Http.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}
	filePath := "./assets/cards/" + strconv.Itoa(id) + ".jpg"
	Http.SendImage(w, filePath)
}
