package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	Cards "ygocarddb/api/cards"
	Users "ygocarddb/api/users"
	Database "ygocarddb/database"
	Middleware "ygocarddb/middlewares"
	models "ygocarddb/models"
)

func Migrate() {
	Database.Instance.AutoMigrate(
		&models.Deck{},
		&models.CardImage{},
		&models.Card{},
		&models.User{})

	log.Println("Database Migration Completed...")
}

func main() {
	r := mux.NewRouter()

	r.Use(Middleware.LoggingMiddleware)
	r.Use(Middleware.ContentTypeApplicationJsonMiddleware)

	Database.Connect()
	Migrate()

	Cards.RegisterRoutes(r)
	Users.RegisterRoutes(r)

	log.Println("Le serveur écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
