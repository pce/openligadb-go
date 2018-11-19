package main

import (
	"fmt"
	"github.com/pce/go-openligadb/openligadb"
)

func main() {
	client := openligadb.NewClient(nil)

	// Fetch matches
	league := "bl1"
	year := 2018
	month := 11

	matches, err := client.GetMatches(league, year, month)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("-- GetMatches --")
	for i, m := range *matches {
		fmt.Printf("%d) MatchID: %d ", (i + 1), m.MatchID)
		fmt.Printf("%v\n", m.MatchDateTime)
		fmt.Printf("%s - %s\n", m.Team1.TeamName, m.Team2.TeamName)
		fmt.Printf("Result: %d:%d\n", m.MatchResults[0].PointsTeam1, m.MatchResults[0].PointsTeam2)
		fmt.Println("-------------------------------------")
	}
	matchID := 51211
	match, err := client.GetMatchByMatchID(matchID)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("-- GGetMatchByMatchID --\n%v\n", match)

}
