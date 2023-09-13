package model

type Person struct {
	ID   int    `json:"ID,omitempty"`
	Name string `json:"name,omitempty"`
}
