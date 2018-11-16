package openligadb

import (
	"cloud.google.com/go/civil"

	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"strings"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://www.openligadb.de/"
	userAgent      = "go-openligadb"

//	wsdlURL = "https://www.openligadb.de/Webservices/Sportsdata.asmx?WSDL"
)

type Client struct {
	UserAgent string
	BaseURL   *url.URL
	client    *http.Client
}

type Team struct {
	TeamId        int
	TeamName      string
	ShortName     string
	TeamIconUrl   string
	TeamGroupName string
}

type MatchResult struct {
	ResultID          int
	ResultName        string
	PointsTeam1       int
	PointsTeam2       int
	ResultTypeID      int
	ResultDescription string
}

type Match struct {
	MatchID         int
	TimeZoneID      string
	LeagueId        int
	MatchDateTime   civil.DateTime
	Team1           Team
	Team2           Team
	MatchIsFinished bool
	MatchResults    []MatchResult
	Goals           []Goal
	// Location string
	NumberOfViewers    string
	LastUpdateDateTime string
}

type Goal struct {
	GoalID         int
	ScoreTeam1     int
	ScoreTeam2     int
	MatchMinute    int
	GoalGetterID   int
	GoalGetterName string
	IsPenalty      bool
	IsOwnGoal      bool
	IsOvertime     bool
	Comment        string
}

func NewClient(httpClient *http.Client) *Client {

	baseURL, _ := url.Parse(defaultBaseURL)

	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := c.client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) GetMatches(league string, year int, month int) (*[]Match, error) {
	// api/getmatchdata/bl1/2016/11

	/*
		if league
		league := "bl1"
		year := "2018"
		month := "11"
	*/

	url := fmt.Sprintf(c.BaseURL.String()+"api/getmatchdata/%s/%d/%d", league, year, month)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	byteArr, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	/*
		s := fmt.Sprintf("%s", byteArr)
		fmt.Printf("%s", s)
	*/
	var data []Match
	err = json.Unmarshal(byteArr, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
