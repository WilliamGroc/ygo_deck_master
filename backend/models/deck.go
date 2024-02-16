package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Deck struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `json:"name"`
	Cards     []uint             `json:"cards"`
	CreatedBy primitive.ObjectID `bson:"createdby" json:"createdBy"`
	CreatedAt string             `json:"createdAt"`
	UpdatedAt string             `json:"updatedAt"`
	IsPublic  bool               `json:"isPublic"`
}
