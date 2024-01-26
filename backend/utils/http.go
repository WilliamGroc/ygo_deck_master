package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)

	if params["id"] == "" {
		err := errors.New("no id provided")
		return 0, err
	}

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return 0, err
	}

	return id, nil
}
