package main

import (
	"log"
	"net/http"
	"ygocarddb/server"
)

func main() {
	app := &server.App{}
	app.Run()

	log.Println("Le serveur écoute sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
