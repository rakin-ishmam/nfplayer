package task

import (
	"github.com/rakin-ishmam/nfplayer/model"
	"github.com/rakin-ishmam/nfplayer/resp"
	"github.com/rakin-ishmam/nfplayer/store"
)

// NameOption holds options for filter players
type NameOption struct {
	MaxID     int
	TeamAdder resp.TeamAdder
	Team      store.Team
	WPool     int
	TNames    []string
}

func (nc NameOption) mapTeams() map[string]bool {
	m := make(map[string]bool)

	for _, v := range nc.TNames {
		m[v] = false
	}

	return m
}

type result struct {
	team model.Team
	err  error
}
