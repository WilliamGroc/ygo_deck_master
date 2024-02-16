package main

import (
	"log"
	"net/http"
	"ygocarddb/api/cards"
	"ygocarddb/api/decks"
	"ygocarddb/api/users"
	"ygocarddb/database"
	"ygocarddb/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitMongoDb()

	r := mux.NewRouter()

	r.Use(middlewares.CorsMiddleware)
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.ContentTypeApplicationJsonMiddleware)

	cards.RegisterRoutes(r)
	users.RegisterRoutes(r)
	decks.RegisterRoutes(r)

	log.Println("Le serveur écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
