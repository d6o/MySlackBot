package pokemon

import (
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	url = "https://pokeapi.co/api/v2/pokemon/%s/"
)

//Response TODO
type Response struct {
	Name    string `json:"name"`
	Weight  int    `json:"weight"`
	Sprites struct {
		BackFemale       interface{} `json:"back_female"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		BackDefault      string      `json:"back_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		BackShiny        string      `json:"back_shiny"`
		FrontDefault     string      `json:"front_default"`
		FrontShiny       string      `json:"front_shiny"`
	} `json:"sprites"`
	HeldItems              []interface{} `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Height                 int           `json:"height"`
	IsDefault              bool          `json:"is_default"`
	Species                struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"species"`
	ID             int `json:"id"`
	Order          int `json:"order"`
	BaseExperience int `json:"base_experience"`
	Types          []struct {
		Slot int `json:"slot"`
		Type struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

//Search returns a pokemon
func Search(pokemon string) (*Response, error) {
	body, err := makeRequest(pokemon)
	if err != nil {
		return nil, err
	}

	return transformStringToResponse(body)
}

func makeRequest(pokemon string) (body string, err error) {
	url := fmt.Sprintf(url, pokemon)
	logrus.Infof("Making Request: %s", url)
	req, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	if req.StatusCode != 200 {
		return "", fmt.Errorf("Status Code: %d", req.StatusCode)
	}
	logrus.Infof("Response body: %s", body)
	return body, nil
}

func transformStringToResponse(body string) (resp *Response, err error) {
	err = json.Unmarshal([]byte(body), &resp)
	return
}
