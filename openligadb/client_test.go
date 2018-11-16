package openligadb

import (
	// "fmt"
	// "net/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	var (
		actual   string
		expected string
	)

	actual = c.BaseURL.String()
	expected = defaultBaseURL

	assert.Equal(t, expected, actual, "unexpected default URL")

	actual = c.UserAgent
	expected = userAgent

	assert.Equal(t, expected, actual, "unexpected default UA")

}

/*
func TestClient(t *testing.T) {

	httpClient := MockClient

	client := NewClient(httpClient)

        // Fetch matches
        league := "bl1"
        year := 2018
        month := 11

        matches, err := client.GetMatches(league, year, month)
        if err != nil {
                fmt.Printf("%v", err)
        }

}
*/
