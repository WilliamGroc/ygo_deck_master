package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"ygocarddb/server/authentication"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(url string, item interface{}) {
	response, err := http.Get(url)

	if err != nil {
		log.Panic(err.Error())
	}

	defer response.Body.Close()

	body, errRead := io.ReadAll(response.Body)

	if errRead != nil {
		log.Panic(errRead.Error())
	}

	json.Unmarshal(body, &item)
}

func Pagination(r *http.Request) *options.FindOptions {
	findOptions := options.Find()

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Panic(err)
		pageInt = 1
	}
	skip := (pageInt - 1) * 20
	findOptions.SetSkip(int64(skip)).SetLimit(20)

	return findOptions
}

func SendImage(w http.ResponseWriter, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-Control", "max-age=86400")

	// Renvoyer le fichier image
	io.Copy(w, file)
}

func GetParamId(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")

	if idParam == "" {
		err := errors.New("no id provided")
		return 0, err
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserId(r *http.Request) (primitive.ObjectID, error) {
	token := authentication.ExtractToken(r)

	claims, err := authentication.ParseToken(token)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	fmt.Println(claims["id"])

	id, err := primitive.ObjectIDFromHex(claims["id"].(string))
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}
