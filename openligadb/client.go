package openligadb

import (
	"cloud.google.com/go/civil"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultBaseURL = "https://www.openligadb.de/"
	userAgent      = "go-openligadb"
	//	wsdlURL = "https://www.openligadb.de/Webservices/Sportsdata.asmx?WSDL"
	verbose = false
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
	MatchID            int
	TimeZoneID         string
	LeagueId           int
	MatchDateTime      civil.DateTime
	Team1              Team
	Team2              Team
	MatchIsFinished    bool
	MatchResults       []MatchResult
	Goals              []Goal
	Location           Location
	NumberOfViewers    string
	LastUpdateDateTime string
}

type Location struct {
	LocationCity    string
	LocationID      int
	LocationStadium string
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

func (c *Client) getMatchData(endpoint string) (*[]Match, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL.String(), endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	byteArr, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	if verbose {
		s := fmt.Sprintf("%s", byteArr)
		fmt.Printf("%s", s)
	}
	var data []Match
	err = json.Unmarshal(byteArr, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) getSingleMatchData(endpoint string) (*Match, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL.String(), endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	byteArr, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	if verbose {
		s := fmt.Sprintf("%s", byteArr)
		fmt.Printf("%s", s)
	}
	var data Match
	err = json.Unmarshal(byteArr, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (c *Client) GetMatches(leagueShortcut string, leagueSaison int, matchday int) (*[]Match, error) {
	endpoint := fmt.Sprintf("api/getmatchdata/%s/%d/%d", leagueShortcut, leagueSaison, matchday)
	return c.getMatchData(endpoint)
}

func (c *Client) GetAvailGroups(leagueShortcut string, leagueSaison int) {}

func (c *Client) GetAvailLeagues() {}

func (c *Client) GetAvailLeaguesBySports(sportID int) {}

func (c *Client) GetAvailSports() {

}

func (c *Client) GetCurrentGroup(leagueShortcut string) {
	// getGroupData
	// return c.getJsonData("api/getcurrentgroup/" + leagueShortcut)
}

func (c *Client) GetCurrentGroupOrderID(leagueShortcut string) {}

func (c *Client) GetGoalGettersByLeagueSaison(leagueShortcut string, leagueSaison int) {
	// return c.getGoalsData("api/getgoalgetters/" + leagueShortcut + "/" + leagueSaison)
}

func (c *Client) GetGoalsByLeagueSaison(leagueShortcut string, leagueSaison string) {}

func (c *Client) GetGoalsByMatch(matchID int) (*[]Match, error) {
	// ?
	return c.getMatchData("api/getgoalsbymatch/" + strconv.Itoa(matchID))
}

func (c *Client) GetLastChangeDateByGroupLeagueSaison(groupOrderID int, leagueShortcut string, leagueSaison string) {
}

func (c *Client) GetLastChangeDateByLeagueSaison(leagueShortcut string, leagueSaison int, matchday int) {
	// ?
	// return c.getJsonData("api/getlastchangedate/" + leagueShortcut + "/" + leagueSaison + "/" + matchday )
}

func (c *Client) GetLastMatch(leagueShortcut string) (*[]Match, error) {
	return c.getMatchData("api/getmatchdata/" + leagueShortcut)
}

func (c *Client) GetLastMatchByLeagueTeam(leagueID int, teamID int) (*[]Match, error) {
	return c.getMatchData("api/getmatchdata/" + strconv.Itoa(leagueID) + "/" + strconv.Itoa(teamID))
}

func (c *Client) GetMatchByMatchID(matchID int) (*Match, error) {
	return c.getSingleMatchData("api/getmatchdata/" + strconv.Itoa(matchID))
}

func (c *Client) GetMatchdataByGroupLeagueSaison(groupOrderID int, leagueShortcut string, leagueSaison string) {
}

func (c *Client) GetMatchdataByGroupLeagueSaisonJSON(groupOrderID int, leagueShortcut string, leagueSaison string) {
}

func (c *Client) GetMatchdataByLeagueDateTime(fromDateTime string, toDateTime string, leagueShortcut string) (*[]Match, error) {
	// ?
	return c.getMatchData("api/getmatchdata/" + leagueShortcut + "/" + fromDateTime + "/" + toDateTime)
}

func (c *Client) GetMatchdataByLeagueSaison(leagueShortcut string, leagueSaison int) (*[]Match, error) {
	return c.getMatchData("api/getmatchdata/" + leagueShortcut + "/" + strconv.Itoa(leagueSaison))
}

func (c *Client) GetMatchdataByTeams(teamID1 int, teamID2 int) (*[]Match, error) {
	return c.getMatchData("api/getmatchdata/" + strconv.Itoa(teamID1) + "/" + strconv.Itoa(teamID2))
}

func (c *Client) GetNextMatch(leagaueShortcut string) {

}

func (c *Client) GetNextMatchByLeagueTeam(leagueID int, teamID int) (*[]Match, error) {
	return c.getMatchData("api/getnextmatchbyleagueteam/" + strconv.Itoa(leagueID) + "/" + strconv.Itoa(teamID))
}

func (c *Client) GetTeamsByLeagueSaison(leagaueShortcut string, leagueSaison int) (*[]Match, error) {
	return c.getMatchData("/api/getavailableteams/" + leagaueShortcut + "/" + strconv.Itoa(leagueSaison))
}
