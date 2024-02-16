package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	Username string             `json:"username" gorm:"unique"`
	Email    string             `json:"email" gorm:"unique"`
	Password string             `json:"password"`
	Decks    []Deck             `json:"decks"`
}
