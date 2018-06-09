package api

import "fmt"

func genURL(base string, id int) string {
	return fmt.Sprintf("%s/api/teams/en/%d.json", base, id)
}
