package model

type Destination struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Visitor *int `json:"visitor,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Photos *[]Image `json:"photos,omitempty"`
	
}