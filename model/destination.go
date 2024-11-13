package model

type Destination struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
}