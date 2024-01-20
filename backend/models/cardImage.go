package models

type CardImage struct {
	Id              uint   `json:"id"`
	ImageUrl        string `json:"image_url"`
	ImageUrlSmall   string `json:"image_url_small"`
	ImageUrlCropped string `json:"image_url_cropped"`
	CardId          uint
}
