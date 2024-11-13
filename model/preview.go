package model

type Preview struct {
	ID int `json:"id,omitempty"`
	EventId Events `json:"event,omitempty"`
	Rating float32 `json:"rating,omitempty"`
	Comment string `json:"comment,omitempty"`
}