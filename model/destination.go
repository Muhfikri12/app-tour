package model

type Destination struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Visitor *int `json:"visitor,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	Photos *[]Image `json:"photos,omitempty"`
}

type Location struct {
	ID int `json:"id,omitempty"`
	Destination *Destination `json:"destination,omitempty"`
	Coordinate string `json:"coordinate,omitempty"`
	FirstDescription string `json:"first_description,omitempty"`
	SecondDescription string `json:"second_description,omitempty"`
}