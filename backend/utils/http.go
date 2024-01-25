package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
