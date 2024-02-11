package cards

import (
	Authentication "ygocarddb/authentication"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/cards").Subrouter()

	s.HandleFunc("", GetCards).Methods("GET")
	s.HandleFunc("/{id}", GetCard).Methods("GET")
	s.HandleFunc("/{id}/image", GetCardImage).Methods("GET")
	s.HandleFunc("/{id}/image/big", GetCardImageBig).Methods("GET")

	secure := s.PathPrefix("").Subrouter()
	secure.Use(Authentication.TokenVerifyMiddleWare)
	secure.HandleFunc("/load", LoadCards).Methods("POST")
}
