package models

//Portfolio is the about.json struct
type Portfolio struct {
	Name        string `json:"name"`
	Descripcion string `json:"descripcion"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}
