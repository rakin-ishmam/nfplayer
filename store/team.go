package store

import "github.com/rakin-ishmam/nfplayer/model"

// Team wraps team store operations
type Team interface {
	ByID(int) (*model.Team, error)
	ByName(string) (*model.Team, error)
}
