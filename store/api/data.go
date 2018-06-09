package api

import (
	"github.com/rakin-ishmam/nfplayer/model"
)

type player struct {
	ID      string `json:"id"`
	FName   string `json:"firstName"`
	LName   string `json:"lastName"`
	Age     string `json:"age"`
	Country string `json:"country"`
}

func (p player) toModel() model.Player {
	m := model.Player{}

	m.ID = p.ID
	m.FName = p.FName
	m.LName = p.LName
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
