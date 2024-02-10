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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResponsePaginated struct {
	Total   int64                  `json:"total"`
	Data    []interface{}          `json:"data"`
	Filters map[string]interface{} `json:"filters"`
}

type ColumnList struct {
	List []string `bson:"list"`
}

func getFilters(coll *mongo.Collection, column string) ColumnList {
	// Aggregate collection to get type list
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "list", Value: bson.D{
				{Key: "$addToSet", Value: "$" + column},
			}},
		}}},
	}

	cursorType, err := coll.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Panic(err)
	}

	var list []ColumnList
	if err = cursorType.All(context.TODO(), &list); err != nil {
		log.Panic(err)
	}

	return list[0]
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	findOptions := Http.Pagination(r)

	search := r.URL.Query().Get("search")
	card_type := r.URL.Query().Get("type")
	card_level := r.URL.Query().Get("level")
	card_attribute := r.URL.Query().Get("attribute")

	var cursor *mongo.Cursor
	var err error
	var filter primitive.D = bson.D{{}}

	if search != "" {
		findOptions.SetSort(bson.D{{Key: "name", Value: 1}})

		filter = bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "en.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "fr.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "de.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "it.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "pt.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
				{{Key: "es.name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}}},
			}},
		}
	}

	if card_type != "" {
		filter = append(filter, bson.E{Key: "frametype", Value: card_type})
	}

	if card_attribute != "" {
		filter = append(filter, bson.E{Key: "attribute", Value: card_attribute})
	}

	if card_level != "" {
		level, err := strconv.Atoi(card_level)
		if err != nil {
			log.Panic(err)
		}
		filter = append(filter, bson.E{Key: "level", Value: level})
	}

	cursor, err = coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Panic(err)
	}

	total, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Panic(err)
	}

	var cards []models.Card

	if err = cursor.All(context.TODO(), &cards); err != nil {
		panic(err)
	}

	// Get filters
	types := getFilters(coll, "frametype")
	attributes := getFilters(coll, "attribute")

	// Convert cards to []interface{}
	var data []interface{}

	for _, card := range cards {
		data = append(data, card)
	}

	json.NewEncoder(w).Encode(ResponsePaginated{
		Total: total,
		Data:  data,
		Filters: map[string]interface{}{
			"types":      types.List,
			"attributes": attributes.List,
		},
	})
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
