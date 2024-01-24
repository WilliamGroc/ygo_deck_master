package models

type CardLang struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Card struct {
	Id            uint        `json:"id"`
	Type          string      `json:"Type"`
	FrameType     string      `json:"frameType"`
	Race          string      `json:"race"`
	Atk           int         `json:"atk"`
	Def           int         `json:"def"`
	Level         int         `json:"level"`
	Attribute     string      `json:"attribute"`
	LinkVal       int         `json:"linkVal"`
	YgoprodeckUrl string      `json:"ygoprodeck_url"`
	Images        []CardImage `json:"card_images"`
	Fr            CardLang    `json:"fr"`
	En            CardLang    `json:"en"`
}
