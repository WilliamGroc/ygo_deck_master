package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Cards "ygocarddb/api/cards"
	Users "ygocarddb/api/users"
	Database "ygocarddb/database"
	Middleware "ygocarddb/middlewares"
)

func main() {
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
