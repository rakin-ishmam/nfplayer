package model

import (
	"fmt"
)

// Player represents player bio
type Player struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Country   string `json:"country"`
}

// FullName generate fullname
func (p Player) FullName() string {
	return fmt.Sprintf("%s %s", p.FName, p.LName)
}
