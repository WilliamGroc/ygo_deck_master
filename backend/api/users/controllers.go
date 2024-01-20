package users

import (
	"encoding/json"
	"log"
	"net/http"

	Authentication "ygocarddb/authentication"
	Database "ygocarddb/database"
	models "ygocarddb/models"

	"github.com/gorilla/mux"
)

// Implement login controller
func Login(w http.ResponseWriter, r *http.Request) {
	var userBody models.User
	json.NewDecoder(r.Body).Decode(&userBody)
	
	var user models.User
	result := Database.Instance.First(&user, "username = ?", userBody.Username)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	err := Authentication.VerifyPassword(user.Password, userBody.Password)

	if err != nil {
		log.Fatal(err)
	}

	token, err := Authentication.GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(token)
}

// Implement register controller
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hash, err := Authentication.HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)

	result := Database.Instance.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	Database.Instance.Select("id", "name", "username", "email").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	Database.Instance.Select("id", "name", "username", "email").First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	
	result := Database.Instance.First(&user, params["id"])

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewDecoder(r.Body).Decode(&user)

	result = Database.Instance.Save(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Database.Instance.Delete(&models.User{}, params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}