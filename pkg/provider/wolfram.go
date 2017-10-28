package provider

import (
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	urlWolfram  = "https://api.wolframalpha.com/v1/result?appid=%s&i=%s"
	errorString = "Wolfram|Alpha did not understand your input"
	errorAnswer = "Aiii, não entendi... :("
)

type (
	Wolfram interface {
		Ask(string) (string, error)
	}

	wolfram struct {
		token string
	}
)

func NewWolfram(token string) Wolfram {
	return &wolfram{
		token: token,
	}
}

//Ask returns weather by City name.
func (w *wolfram) Ask(question string) (string, error) {
	body, err := w.makeRequest(question)
	if err != nil {
		return "Error asking question :(", err
	}

	if strings.Contains(body, errorString) {
		body = errorAnswer
	}

	return body, nil
}

func (w *wolfram) makeRequest(question string) (body string, err error) {
	url := fmt.Sprintf(urlWolfram, w.token, question)
	logrus.Infof("Making Request: %s", url)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	logrus.Infof("Response body: %s", body)
	return body, nil
}
