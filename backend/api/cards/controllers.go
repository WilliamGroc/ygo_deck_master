package cards

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	Database "ygocarddb/database"
	models "ygocarddb/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
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
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	var card models.Card
	cursor := coll.FindOne(context.TODO(), bson.D{
		{Key: "id", Value: id},
	})

	if cursor.Err() != nil {
		log.Fatal(cursor.Err())
	}

	cursor.Decode(&card)

	json.NewEncoder(w).Encode(card)
}

// url: https://db.ygoprodeck.com/api/v7/cardinfo.php

type ExternalCardData struct {
	Data    []models.Card `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func LoadCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	response, apiError := http.Get("https://db.ygoprodeck.com/api/v7/cardinfo.php")

	if apiError != nil {
		log.Fatal(apiError)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Panic(err.Error())
	}

	var data ExternalCardData

	json.Unmarshal(body, &data)

	for _, item := range data.Data {
		_, errInsert := coll.InsertOne(context.TODO(), item)

		if errInsert != nil {
			log.Panic(errInsert.Error())
		}
	}
}
