package cards

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	Database "ygocarddb/database"
	models "ygocarddb/models"

	"github.com/gorilla/mux"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	db := Database.Instance

	var cards []models.Card
	result := db.Find(&cards)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(cards)
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance

	var card models.Card
	result := db.First(&card, params["id"])

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(card)
}

func CreateCard(w http.ResponseWriter, r *http.Request) {
	db := Database.Instance

	var card models.Card
	json.NewDecoder(r.Body).Decode(&card)

	result := db.Create(&card)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(card)
}

func UpdateCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance

	var card models.Card
	result := db.First(&card, params["id"])

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewDecoder(r.Body).Decode(&card)

	result = db.Save(&card)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(card)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database.Instance

	var card models.Card
	result := db.Delete(&card, params["id"])

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Fprintf(w, "Card supprimé")
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

	db := Database.Instance

	for _, card := range data.Data {
		result := db.Create(&card)

		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}

}
