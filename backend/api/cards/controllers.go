package cards

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	Database "ygocarddb/database"
	models "ygocarddb/models"
	Http "ygocarddb/utils"

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

type JSONCardFile struct {
	Data []models.Card `json:"data"`
}

type JSONIDMapping map[string][]int

func LoadCards(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("Card")

	content, err := os.ReadFile("../card_list.json")
	if err != nil {
		log.Panic(err.Error())
	}
	var data JSONCardFile
	json.Unmarshal(content, &data)

	contentId, errId := os.ReadFile("../en_id_mapping.json")
	if errId != nil {
		log.Panic(errId.Error())
	}
	var dataId JSONIDMapping
	json.Unmarshal(contentId, &dataId)

	fmt.Println()

	for _, item := range data.Data {

		cardId := dataId[item.Name][0]

		fmt.Println(item.Name, cardId)

		Http.Get("https://db.ygorganization.com/data/card/"+strconv.FormatInt(int64(cardId), 10), &item)

		_, errInsert := coll.InsertOne(context.TODO(), item)

		time.Sleep(2 * time.Microsecond)

		if errInsert != nil {
			log.Panic(errInsert.Error())
		}
	}
}
