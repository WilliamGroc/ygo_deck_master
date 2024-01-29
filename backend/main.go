package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	Cards "ygocarddb/api/cards"
	Users "ygocarddb/api/users"
	Database "ygocarddb/database"
	Middleware "ygocarddb/middlewares"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Database.InitMongoDb()

	r := mux.NewRouter()

	r.Use(Middleware.CorsMiddleware)
	r.Use(Middleware.LoggingMiddleware)
	r.Use(Middleware.ContentTypeApplicationJsonMiddleware)

	Cards.RegisterRoutes(r)
	Users.RegisterRoutes(r)

	log.Println("Le serveur écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
