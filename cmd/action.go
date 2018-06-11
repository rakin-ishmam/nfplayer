package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rakin-ishmam/nfplayer/store/api"

	"github.com/rakin-ishmam/nfplayer/resp"
	"github.com/rakin-ishmam/nfplayer/store"
	"github.com/rakin-ishmam/nfplayer/task"
)

func apiStore() *store.Stores {
	return &store.Stores{
		Team: api.NewTeam("https://vintagemonster.onefootball.com"),
	}
}

func byid(s *store.Stores, id int) {
	t, err := s.Team.ByID(id)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	resp := resp.NewRGenerator()
	resp.TeamAdd(*t)
	resp.Print(os.Stdout)

	// rt := resp.NewTeam(os.Stdout, t)
	// rt.PrtPlayers()
}

func byname(s *store.Stores, name string, maxid, wpool int) {

	resp := resp.NewRGenerator()
	fd := task.NameFinder(task.NameOption{
		MaxID:     maxid,
		TeamAdder: resp,
		Team:      s.Team,
		WPool:     wpool,
		TNames:    []string{name},
	})

	err := fd.Find()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Print(os.Stdout)
}

func def(s *store.Stores, maxid, wpool int) {
	log.Println("maxid and wpool", maxid, wpool)

	teams := []string{
		"Germany",
		"England",
		"France",
		"Spain",
		"Manchester Utd",
		"Arsenal",
		"Chelsea",
		"Barcelona",
		"Real Madrid",
		"FC Bayern Munich",
	}

	resp := resp.NewRGenerator()
	fd := task.NameFinder(task.NameOption{
		MaxID:     maxid,
		TeamAdder: resp,
		Team:      s.Team,
		WPool:     wpool,
		TNames:    teams,
	})

	err := fd.Find()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Print(os.Stdout)

}
