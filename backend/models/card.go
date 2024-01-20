package models

type Card struct {
	Id            uint        `json:"id"`
	Name          string      `json:"name"`
	Type          string      `json:"Type"`
	FrameType     string      `json:"frameType"`
	Desc          string      `json:"desc"`
	Race          string      `json:"race"`
	Atk           int         `json:"atk"`
	Def           int         `json:"def"`
	Level         int         `json:"level"`
	Attribute     string      `json:"attribute"`
	LinkVal       int         `json:"linkVal"`
	YgoprodeckUrl string      `json:"ygoprodeck_url"`
	Images        []CardImage `json:"card_images"`
}
