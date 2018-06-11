package task

import (
	"context"
	"log"
)

// nameFind finds player list by country name
type nameFind struct {
	option NameOption
}

func (nf *nameFind) genID(lb *loadBalancer) {
	for i := 1; i <= nf.option.MaxID; i++ {
		lb.req(i)
	}
	lb.closeReq()
}

// Find runs filter operation
func (nf *nameFind) Find() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lb := newLoadBalan(ctx, nf.option.Team, nf.option.WPool)
	c := lb.run()

	go nf.genID(lb)

	mTeam := nf.option.mapTeams()
	var err error
	totFound := 0

	for res := range c {
		if err != nil {
			continue
		}

		if res.err != nil {
			err = res.err
			cancel()
			continue
		}

		if exist, ok := mTeam[res.team.Name]; ok && !exist {
			totFound++
			mTeam[res.team.Name] = true
			nf.option.TeamAdder.TeamAdd(res.team)
		}

		if totFound == len(mTeam) {
			cancel()
		}

	}

	log.Println("tot found", totFound)

	return err
}

// NameFinder returns intance to filter player by team name
func NameFinder(option NameOption) Finder {
	return &nameFind{
		option: option,
	}
}
