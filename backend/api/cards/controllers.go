package cards

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	Database "ygocarddb/database"
	models "ygocarddb/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	cursor, err := coll.Find(context.TODO(), bson.D{{}})

	var cards []models.Card

	if err = cursor.All(context.TODO(), &cards); err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(cards)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	cursor, err := coll.FindOne(context.TODO(), bson.D{{}})

	var card models.Card
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

	var data ExternalCardData

	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Panic(err.Error())
	}

	for _, item := range data.Data {
		fmt.Println(item.Name)
		_, err := coll.InsertOne(context.TODO(), item)

		if err != nil {
			log.Panic(err.Error())
		}
	}
}
