package users

import (
	"encoding/json"
	"log"
	"net/http"
	"ygocarddb/authentication"
	"ygocarddb/database"
	"ygocarddb/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Implement login controller
func Login(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("User")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var result models.User
	err := coll.FindOne(r.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	err = authentication.VerifyPassword(result.Password, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Password is incorrect")
		return
	}

	token, err := authentication.GenerateToken(result)

	if err != nil {
		log.Panic(err)
	}

	response := map[string]string{
		"token": token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Implement register controller
func Register(w http.ResponseWriter, r *http.Request) {
	db := database.MongoInstance
	coll := db.Collection("User")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, err := authentication.HashPassword(user.Password)

	if err != nil {
		log.Panic(err)
	}

	user.Password = string(hash)

	result, error := coll.InsertOne(r.Context(), user)

	if error != nil {
		log.Panic(error)
	}

	coll.FindOne(r.Context(), bson.D{{Key: "_id", Value: result.InsertedID}}).Decode(&user)

	token, err := authentication.GenerateToken(user)

	if err != nil {
		log.Panic(err)
	}

	response := map[string]string{
		"token": token,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// var user models.User
	// database.Instance.Select("id", "name", "username", "email").First(&user, params["id"])
	// json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// var user models.User

	// result := database.Instance.First(&user, params["id"])

	// if result.Error != nil {
	// 	log.Panic(result.Error)
	// }

	// json.NewDecoder(r.Body).Decode(&user)

	// result = database.Instance.Save(&user)

	// if result.Error != nil {
	// 	log.Panic(result.Error)
	// }

	// json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// database.Instance.Delete(&models.User{}, params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}
