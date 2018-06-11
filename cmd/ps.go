// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"
	"log"
	"math"

	"github.com/spf13/cobra"
)

type flags struct {
	TeamName string
	TeamID   int
	Logg     bool

	MaxID int
	WPool int
}

var fls flags

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "players",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		if !fls.Logg {
			log.SetOutput(ioutil.Discard)
		}

		store := apiStore()

		if fls.TeamID > -math.MaxInt32 {
			byid(store, fls.TeamID)
			return
		}

		if len(fls.TeamName) > 0 {
			byname(store, fls.TeamName, fls.MaxID, fls.WPool)
			return
		}
		def(store, fls.MaxID, fls.WPool)

	},
}

func init() {
	psCmd.Flags().StringVarP(&fls.TeamName, "tname", "n", "", "get players by team name")
	psCmd.Flags().IntVarP(&fls.TeamID, "tid", "i", -math.MaxInt32, "get players by team id")
	psCmd.Flags().BoolVarP(&fls.Logg, "log", "l", false, "show log")

	psCmd.Flags().IntVarP(&fls.MaxID, "maxid", "m", 100, "max id range")
	psCmd.Flags().IntVarP(&fls.WPool, "wpool", "w", 40, "worker pool")

	rootCmd.AddCommand(psCmd)
}
