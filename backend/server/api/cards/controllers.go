package cards

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"ygocarddb/models"
	"ygocarddb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (t *CardsRoutes) GetCards(w http.ResponseWriter, r *http.Request) {
	coll := t.DB.Collection("Card")

	findOptions := utils.Pagination(r)

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

	cursor, err = coll.Find(r.Context(), filter, findOptions)
	if err != nil {
		log.Panic(err)
	}

	total, err := coll.CountDocuments(r.Context(), filter)
	if err != nil {
		log.Panic(err)
	}

	var cards []models.Card

	if err = cursor.All(r.Context(), &cards); err != nil {
		panic(err)
	}

	// Get filters
	types := utils.GetFilters(coll, "frametype", r)
	attributes := utils.GetFilters(coll, "attribute", r)

	// Convert cards to []interface{}
	var data []interface{}

	for _, card := range cards {
		data = append(data, card)
	}

	json.NewEncoder(w).Encode(utils.ResponsePaginated{
		Total: total,
		Data:  data,
		Filters: map[string]interface{}{
			"types":      types.List,
			"attributes": attributes.List,
		},
	})
}

func (t *CardsRoutes) GetCard(w http.ResponseWriter, r *http.Request) {
	coll := t.DB.Collection("Card")

	id, err := utils.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}

	var card models.Card
	cursor := coll.FindOne(r.Context(), bson.D{
		{Key: "id", Value: id},
	})

	if cursor.Err() != nil {
		log.Panic(cursor.Err())
	}

	cursor.Decode(&card)

	json.NewEncoder(w).Encode(card)
}

func (t *CardsRoutes) GetCardImage(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}
	filePath := "./assets/cards_small/" + strconv.Itoa(id) + ".jpg"
	utils.SendImage(w, filePath)
}

func (t *CardsRoutes) GetCardImageBig(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamId(r)
	if err != nil {
		log.Panic(err)
	}
	filePath := "./assets/cards/" + strconv.Itoa(id) + ".jpg"
	utils.SendImage(w, filePath)
}
