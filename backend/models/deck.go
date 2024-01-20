package models

import "gorm.io/gorm"

type Deck struct {
	gorm.Model
	Name   string `json:"name"`
	Cards  []Card `gorm:"many2many:cards;json:"cards"`
	UserId uint   `json:"user_id"`
}
