package rtm

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
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
	body, err := makeRequest(token)
	if err != nil {
		return nil, err
	}

	return transformStringToResponse(body)
}

func makeRequest(token string) (body string, err error) {
	url := fmt.Sprintf(url, token)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	return body, nil
}

func transformStringToResponse(body string) (resp *Response, err error) {
	err = json.Unmarshal([]byte(body), &resp)
	return
}
