package users

import (
	Authentication "ygocarddb/authentication"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/users").Subrouter()

	s.HandleFunc("/login", Login).Methods("POST")
	s.HandleFunc("/register", Register).Methods("POST")

	secure := s.PathPrefix("").Subrouter()
	secure.Use(Authentication.TokenVerifyMiddleWare)

	secure.HandleFunc("", GetUsers).Methods("GET")
	secure.HandleFunc("/{id}", GetUser).Methods("GET")
	secure.HandleFunc("/{id}", UpdateUser).Methods("PUT")
	secure.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
}