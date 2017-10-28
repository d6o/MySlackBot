package rtm

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	url = "https://slack.com/api/rtm.start?token=%s"
)

//Response has some attributes returned by api/rtm.start.
type Response struct {
	Ok   bool `json:"ok"`
	Self struct {
		ID string `json:"id"`
	} `json:"self"`
	URL string `json:"url"`
}

//New returns a RTM response object.
func New(token string) (*Response, error) {
	resp, err := makeRequest(token)
	if err != nil {
		return nil, err
	}

	return transformStringToResponse(resp)
}

func makeRequest(token string) (*http.Response, error) {
	url := fmt.Sprintf(url, token)
	logrus.Infof("Making Request: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func transformStringToResponse(response *http.Response) (*Response, error) {
	resp := Response{}
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
