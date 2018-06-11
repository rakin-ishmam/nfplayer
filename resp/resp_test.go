package resp_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/rakin-ishmam/nfplayer/model"

	"github.com/rakin-ishmam/nfplayer/resp"
)

func TestResp(t *testing.T) {
	g := resp.NewRGenerator()

	addTeam(g)
	check(t, g)
}

func addTeam(ta resp.TeamAdder) {
	team := model.Team{
		ID:         1,
		Name:       "test",
		IsNational: false,
		Players:    players(),
	}
	ta.TeamAdd(team)
}

func players() []model.Player {
	pp := []model.Player{}

	pp = append(pp, model.Player{
		ID:      "player1",
		FName:   "player1",
		LName:   "player1 l",
		Age:     "10",
		Country: "country 1",
	})

	pp = append(pp, model.Player{
		ID:      "player2",
		FName:   "player2",
		LName:   "player2 l",
		Age:     "10",
		Country: "country 2",
	})

	return pp
}

func check(t *testing.T, ta resp.Printer) {
	bs := []byte{}
	buf := bytes.NewBuffer(bs)
	ta.Print(buf)

	exp := fmt.Sprintf("1. player1 player1 l; 10; country 1, test\n2. player2 player2 l; 10; country 2, test\n")
	res := buf.String()

	if res != exp {
		t.Errorf("expected result is \n%vbut got \n%vdd", exp, res)
	}

}
