package api

import (
	"github.com/rakin-ishmam/nfplayer/model"
)

type player struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Country   string `json:"country"`
}

func (p player) toModel() model.Player {
	m := model.Player{}

	m.ID = p.ID
	m.FirstName = p.FirstName
	m.LastName = p.LastName
	m.Age = p.Age
	m.Country = p.Country

	return m
}

type team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []player `json:"players"`
}

func (t team) toModel() model.Team {
	m := model.Team{}

	m.ID = t.ID
	m.Name = t.Name
	for _, p := range t.Players {
		m.Players = append(m.Players, p.toModel())
	}

	return m
}

type data struct {
	Team team `json:"team"`
}

type resp struct {
	Data data `json:"data"`
}
