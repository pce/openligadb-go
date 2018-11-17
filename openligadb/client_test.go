package openligadb

import (
	"bytes"
	"net/http"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"strings"
)


type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}

func TestClientMockRequest(t *testing.T) {


	b, _ := ioutil.ReadFile("../testdata/fixtures/matches.json")

	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// do whatever you want
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader(b)),
			}, nil
		},
	}

	request, _ := http.NewRequest("GET", "https://www.example.com", nil)
	response, _ := client.Do(request)
	if response.StatusCode == http.StatusBadRequest {
		t.Error("invalid response status code")
	}


	assert.NotEmpty(t, response, "empty response")
	
	responseBytes, _ := ioutil.ReadAll(response.Body)


	expectedFirst := "["
	expectedLast := "]"
	actual := string(responseBytes)

	actual = strings.TrimSpace(actual)

        firstChar := actual[:1]
        lastChar := actual[len(actual)-1:]

	assert.Equal(t, expectedFirst, firstChar, "unexpected response")
	assert.Equal(t, expectedLast, lastChar, "unexpected response")

}


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

