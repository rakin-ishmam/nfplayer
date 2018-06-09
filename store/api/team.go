package api

import (
	"github.com/rakin-ishmam/nfplayer/model"
	"github.com/rakin-ishmam/nfplayer/store"
)

type team struct {
	baseURL string
}

func (t *team) ByID(id int) (*model.Team, error) {
	return nil, nil
}

func (t *team) ByName(name string) (*model.Team, error) {
	return nil, nil
}

// New returns instace of store.Team for api driver
func New(baseURL string) store.Team {
	return &team{
		baseURL: baseURL,
	}
}
