package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/rakin-ishmam/nfplayer/model"
)

// fetch and parse
type fnp struct {
	URL string

	team model.Team

	hresp *http.Response

	err error
}

func (f *fnp) get() {
	if f.hasErr() {
		return
	}

	f.hresp, f.err = http.Get(f.URL)
}

func (f *fnp) parse() {
	if f.hasErr() {
		return
	}

	resp := resp{}

	if err := json.NewDecoder(f.hresp.Body).Decode(&resp); err != nil {
		f.err = err
		return
	}

	if resp.isNotFound() {
		f.err = fmt.Errorf("%s", resp.Message)
		return
	}

	f.team = resp.Data.Team.toModel()
}

func (f *fnp) resp() (*model.Team, error) {
	if f.hasErr() {
		return nil, f.err
	}

	return &f.team, nil
}

func (f *fnp) hasErr() bool {
	if f.err != nil {
		return true
	}

	return false
}

// Team handles api team store
type Team struct {
	baseURL string
}

// ByID returns team by id
func (t *Team) ByID(id int) (*model.Team, error) {
	url := genURL(t.baseURL, id)

	job := fnp{
		URL: url,
	}

	job.get()
	job.parse()
	return job.resp()
}

// ByName returns team by name
func (t *Team) ByName(name string) (*model.Team, error) {
	lname := strings.ToLower(name)

	for i := 0; i < math.MaxInt32; i++ {
		m, err := t.ByID(i)
		if err != nil {
			return nil, err
		}

		if strings.ToLower(m.Name) == lname {
			return m, nil
		}
	}
	return nil, nil
}

// NewTeam returns instace of store.Team for api driver
func NewTeam(baseURL string) *Team {
	return &Team{
		baseURL: baseURL,
	}
}
