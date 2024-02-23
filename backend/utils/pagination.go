package utils

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
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

func GetFilters(coll *mongo.Collection, column string, r *http.Request) ColumnList {
	// Aggregate collection to get type list
	pipeline := mongo.Pipeline{
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "list", Value: bson.D{
				{Key: "$addToSet", Value: "$" + column},
			}},
		}}},
	}

	cursorType, err := coll.Aggregate(r.Context(), pipeline)
	if err != nil {
		log.Panic(err)
	}

	var list []ColumnList
	if err = cursorType.All(r.Context(), &list); err != nil {
		log.Panic(err)
	}

	return list[0]
}
