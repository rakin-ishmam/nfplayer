package resp

import (
	"fmt"
	"io"

	"github.com/emirpasic/gods/trees/btree"
	"github.com/rakin-ishmam/nfplayer/model"
)

type pstore struct {
	players *btree.Tree
}

func (p *pstore) get(mp model.Player) player {
	pre, ok := p.players.Get(mp.FullName())
	if ok {
		return pre.(player)
	}

	np := player{
		teams: btree.NewWithStringComparator(3),
	}
	np.fromModel(mp)

	return np
}

func (p *pstore) TeamAdd(team model.Team) {
	for _, mplayer := range team.Players {
		pl := p.get(mplayer)
		pl.AddTeam(team)

		p.players.Put(mplayer.FullName(), pl)

	}
}

func (p pstore) Print(wr io.Writer) {
	it := p.players.Iterator()
	for i := 1; it.Next(); i++ {
		pl := it.Value().(player)
		fmt.Fprintf(wr, "%d. %s; %s; %s\n", i, pl.fullName, pl.age, pl.strTeams())
	}
}

// player represents player
type player struct {
	fullName string
	country  string
	age      string
	teams    *btree.Tree
}

func (p *player) fromModel(m model.Player) {
	p.country = m.Country
	p.age = m.Age
	p.fullName = m.FullName()
}

func (p *player) AddTeam(t model.Team) {
	if t.IsNational {
		return
	}

	p.teams.Put(t.Name, true)
}

func (p *player) strTeams() string {
	team := p.country
	it := p.teams.Iterator()
	for it.Next() {
		team = fmt.Sprintf("%s, %s", team, it.Key().(string))
	}

	return team
}

// NewRGenerator returns instance of RespGenerator
func NewRGenerator() RepsGenerator {
	return &pstore{
		players: btree.NewWithStringComparator(3),
	}
}
