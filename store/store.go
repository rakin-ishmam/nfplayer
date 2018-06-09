package store

import "github.com/rakin-ishmam/nfplayer/store/api"

// Stores holds all kind of stores
type Stores struct {
	Team Team
}

// NewAPI returns api store
func NewAPI(baseURL string) *Stores {
	return &Stores{
		Team: api.NewTeam(baseURL),
	}
}
