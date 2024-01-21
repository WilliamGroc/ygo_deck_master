package models

type Deck struct {
	Name   string `json:"name"`
	Cards  []Card `json:"cards"`
	UserId uint   `json:"user_id"`
}
