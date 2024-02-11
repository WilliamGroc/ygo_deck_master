package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Authentication "ygocarddb/authentication"
	Database "ygocarddb/database"
	models "ygocarddb/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Implement login controller
func Login(w http.ResponseWriter, r *http.Request) {
	db := Database.MongoInstance
	coll := db.Collection("User")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user)

	var result models.User
	err := coll.FindOne(r.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	err = Authentication.VerifyPassword(result.Password, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Password is incorrect")
		return
	}

	token, err := Authentication.GenerateToken(result)

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
	db := Database.MongoInstance
	coll := db.Collection("User")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, err := Authentication.HashPassword(user.Password)

	if err != nil {
		log.Panic(err)
	}

	user.Password = string(hash)

	_, error := coll.InsertOne(r.Context(), user)

	if error != nil {
		log.Panic(error)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(true)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// var user models.User
	// Database.Instance.Select("id", "name", "username", "email").First(&user, params["id"])
	// json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// var user models.User

	// result := Database.Instance.First(&user, params["id"])

	// if result.Error != nil {
	// 	log.Panic(result.Error)
	// }

	// json.NewDecoder(r.Body).Decode(&user)

	// result = Database.Instance.Save(&user)

	// if result.Error != nil {
	// 	log.Panic(result.Error)
	// }

	// json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)

	// Database.Instance.Delete(&models.User{}, params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}
