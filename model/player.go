package model

// Player represents player bio
type Player struct {
	ID      string `json:"id"`
	FName   string `json:"firstName"`
	LName   string `json:"lastName"`
	Age     string `json:"age"`
	Country string `json:"country"`
}
