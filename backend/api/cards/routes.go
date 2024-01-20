package cards

import (
	// Authentication "ygocarddb/authentication"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	s := r.PathPrefix("/cards").Subrouter()

	// s.Use(Authentication.TokenVerifyMiddleWare)
	s.HandleFunc("/load", LoadCards).Methods("POST")
	s.HandleFunc("", GetCards).Methods("GET")
	s.HandleFunc("/{id}", GetCard).Methods("GET")
	s.HandleFunc("", CreateCard).Methods("POST")
	s.HandleFunc("/{id}", UpdateCard).Methods("PUT")
	s.HandleFunc("/{id}", DeleteCard).Methods("DELETE")
}
