package wolfram

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	url = "https://api.wolframalpha.com/v1/result?appid=%s&i=%s"
)

//Short TODO
type Short struct {
	token string
}

//New TODO
func New(token string) *Short {
	return &Short{
		token: token,
	}
}

//ByName returns weather by City name.
func (w *Short) Ask(question string) (string, error) {
	body, err := w.makeRequest(question)
	if err != nil {
		return "Error asking question :(", err
	}

	return body, nil
}

func (w *Short) makeRequest(question string) (body string, err error) {
	url := fmt.Sprintf(url, w.token, question)
	logrus.Infof("Making Request: %s", url)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	logrus.Infof("Response body: %s", body)
	return body, nil
}
