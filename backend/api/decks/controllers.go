package decks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"ygocarddb/database"
	"ygocarddb/models"
	"ygocarddb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListDecks(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("Deck")

	var filters = bson.D{{
		Key: "$or",
		Value: bson.A{
			bson.D{{Key: "ispublic", Value: true}},
			bson.D{{Key: "createdby", Value: r.URL.Query().Get("userId")}},
		}}, {
		Key:   "name",
		Value: bson.D{{Key: "$regex", Value: r.URL.Query().Get("search")}},
	}}

	var decks []models.Deck
	cursor, err := coll.Find(r.Context(), filters)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error fetching decks")
		return
	}

	if err = cursor.All(r.Context(), &decks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error decoding decks")
		return
	}

	total, err := coll.CountDocuments(r.Context(), filters)
	if err != nil {
		log.Panic(err)
	}

	var data []interface{}

	for _, card := range decks {
		data = append(data, card)
	}

	json.NewEncoder(w).Encode(utils.ResponsePaginated{
		Total:   total,
		Data:    data,
		Filters: map[string]interface{}{},
	})
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("Deck")

	id, err := utils.GetParamId(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid id")
		return
	}

	cursor := coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: id}})

	if cursor.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Deck not found")
		return
	}

	var deck models.Deck
	err = cursor.Decode(&deck)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error decoding deck")
		return
	}

	json.NewEncoder(w).Encode(deck)
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("Deck")

	userId, err := utils.GetUserId(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting user id")
		return
	}

	type Input struct {
		Name     string `json:"name"`
		Cards    []uint `json:"cards"`
		IsPublic bool   `json:"isPublic"`
	}

	var input Input
	json.NewDecoder(r.Body).Decode(&input)

	deck := models.Deck{
		Id:        primitive.NewObjectID(),
		Name:      input.Name,
		Cards:     input.Cards,
		CreatedBy: userId,
		CreatedAt: time.Now().Format(time.RFC3339),
		IsPublic:  input.IsPublic,
	}

	result, err := coll.InsertOne(r.Context(), deck)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error creating deck")
		return
	}

	var output models.Deck
	err = coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&output)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error decoding deck")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func UpdateDeck(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("Deck")

	id, err := utils.GetParamId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	userId, err := utils.GetUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting user id")
		return
	}

	var deck models.Deck
	err = coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: id}}).Decode(&deck)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Deck not found")
		return
	}

	if userId != deck.CreatedBy {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You can't update this deck")
		return
	}

	var input models.Deck
	json.NewDecoder(r.Body).Decode(&input)

	_, err = coll.UpdateOne(r.Context(), bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: input}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error updating deck")
		return
	}

	var output models.Deck
	err = coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: id}}).Decode(&output)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error decoding deck")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func DeleteDeck(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("Deck")

	id, err := utils.GetParamId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	userId, err := utils.GetUserId(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting user id")
		return
	}

	var deck models.Deck
	err = coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: id}}).Decode(&deck)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Deck not found")
		return
	}

	if userId != deck.CreatedBy {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("You can't delete this deck")
		return
	}

	_, err = coll.DeleteOne(r.Context(), bson.D{{Key: "_id", Value: id}})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error deleting deck")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deck deleted")
}
