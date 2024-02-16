package decks

import (
	"ygocarddb/authentication"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/decks").Subrouter()

	s.HandleFunc("", ListDecks).Methods("GET")
	s.HandleFunc("/{id}", GetDeck).Methods("GET")

	secure := s.PathPrefix("").Subrouter()
	secure.Use(authentication.TokenVerifyMiddleWare)

	secure.HandleFunc("", CreateDeck).Methods("POST")
	secure.HandleFunc("/{id}", UpdateDeck).Methods("PUT")
	secure.HandleFunc("/{id}", DeleteDeck).Methods("DELETE")
}
