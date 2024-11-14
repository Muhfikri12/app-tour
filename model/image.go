package model

type Image struct {
	Image string `json:"image_url,omitempty"`
	DestinationId int `json:"destination_id,omitempty"`
	Description string `json:"description,omitempty"`
}