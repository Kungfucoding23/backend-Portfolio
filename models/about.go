package models

//About is the about.json struct
type About struct {
	Name       string   `json:"name"`
	Profession string   `json:"profession"`
	Photo      string   `json:"photo"`
	AboutMe    string   `json:"about_me"`
	Skills     []string `json:"skills"`
}
