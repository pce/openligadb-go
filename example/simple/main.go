package main
 
import (
	"github.com/pce/go-openligadb/openligadb"
	"fmt"
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

	for i, m := range *matches {
		fmt.Printf("%v. %v\n", i+1, m)
	}
	// fmt.Printf("%v", matches)
}

