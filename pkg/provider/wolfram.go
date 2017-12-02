package provider

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	urlWolfram = "http://api.wolframalpha.com/v2/query?appid=%s&input=%s&output=JSON&format=image,plaintext"
)

type (
	Wolfram interface {
		Ask(string) (*WolframResponse, error)
	}

	wolfram struct {
		token string
	}

	WolframResponse struct {
		Queryresult struct {
			Success       bool    `json:"success"`
			Error         bool    `json:"error"`
			Pods          []struct {
				Title string `json:"title"`
				Subpods    []struct {
					Title string `json:"title"`
					Plaintext string `json:"plaintext"`
				} `json:"subpods"`
			} `json:"pods"`
			Didyoumeans   []struct {
				Score string `json:"score"`
				Level string `json:"level"`
				Val   string `json:"val"`
			} `json:"didyoumeans"`
		} `json:"queryresult"`
	}
)

func NewWolfram(token string) Wolfram {
	return &wolfram{
		token: token,
	}
}

//Ask returns weather by City name.
func (w *wolfram) Ask(question string) (*WolframResponse, error) {
	body, err := w.makeRequest(question)
	if err != nil {
		return nil, err
	}

	return w.transformStringToResponse(body)
}

func (w *wolfram) makeRequest(question string) (body string, err error) {
	url := fmt.Sprintf(urlWolfram, w.token, question)
	logrus.Infof("Making Request: %s", url)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}

	return body, nil
}

func (w *wolfram) transformStringToResponse(body string) (*WolframResponse, error) {
	resp := WolframResponse{}
	fmt.Println(body)
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
