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
	"ygocarddb/database"
	"ygocarddb/models"
	"ygocarddb/utils"
)

// url: https://db.ygoprodeck.com/api/v7/cardinfo.php

type CardFromApi struct {
	Id            uint               `json:"id"`
	Name          string             `json:"name"`
	Type          string             `json:"Type"`
	FrameType     string             `json:"frameType"`
	Race          string             `json:"race"`
	Atk           int                `json:"atk"`
	Def           int                `json:"def"`
	Level         int                `json:"level"`
	Attribute     string             `json:"attribute"`
	LinkVal       int                `json:"linkVal"`
	YgoprodeckUrl string             `json:"ygoprodeck_url"`
	Images        []models.CardImage `json:"card_images"`
	Desc          string             `json:"desc"`
}

type JSONCardFile struct {
	Data []CardFromApi `json:"data"`
}

type JSONIDMapping map[string][]int

type CardData struct {
	CardData models.Card `json:"cardData"`
}

func LoadCards(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
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

	cardsLen := len(data.Data)

	for i, item := range data.Data {
		var cardToInsert models.Card = models.Card{
			Id:            uint(item.Id),
			Name:          item.Name,
			Type:          item.Type,
			FrameType:     item.FrameType,
			Images:        item.Images,
			Race:          item.Race,
			Atk:           item.Atk,
			Def:           item.Def,
			Level:         item.Level,
			Attribute:     item.Attribute,
			LinkVal:       item.LinkVal,
			YgoprodeckUrl: item.YgoprodeckUrl,
		}

		if len(dataId[item.Name]) == 0 {
			cardToInsert.En.Name = item.Name
			cardToInsert.En.EffectText = item.Desc
		} else {
			cardId := dataId[item.Name][0]

			var cardWithTranslation CardData

			utils.Get("https://db.ygorganization.com/data/card/"+strconv.FormatInt(int64(cardId), 10), &cardWithTranslation)

			cardToInsert.Fr = cardWithTranslation.CardData.Fr
			cardToInsert.En = cardWithTranslation.CardData.En
			cardToInsert.Es = cardWithTranslation.CardData.Es
			cardToInsert.It = cardWithTranslation.CardData.It
			cardToInsert.De = cardWithTranslation.CardData.De
			cardToInsert.Ja = cardWithTranslation.CardData.Ja
			cardToInsert.Ko = cardWithTranslation.CardData.Ko
			cardToInsert.Pt = cardWithTranslation.CardData.Pt
		}

		fmt.Println(i, "/", cardsLen, cardToInsert.Name)
		_, errInsert := coll.InsertOne(context.TODO(), cardToInsert)

		time.Sleep(1 * time.Microsecond)

		if errInsert != nil {
			log.Panic(errInsert.Error())
		}
	}
}
