package task

import (
	"context"
	"fmt"
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
	if err != nil {
		return nil
	}

	return nf.checkAllTeam(mTeam)
}

func (nf nameFind) checkAllTeam(teams map[string]bool) error {
	notFound := ""
	for k := range teams {
		if found := teams[k]; !found {
			if len(notFound) > 0 {
				notFound = fmt.Sprintf("%s, %s", notFound, k)
				continue
			}

			notFound = k
		}
	}

	if len(notFound) > 0 {
		return fmt.Errorf("not found: %s", notFound)
	}

	return nil
}

// NameFinder returns intance to filter player by team name
func NameFinder(option NameOption) Finder {
	return &nameFind{
		option: option,
	}
}
