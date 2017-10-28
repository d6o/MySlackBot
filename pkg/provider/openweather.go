package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	openWeatherURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"
)

type (
	Weather interface {
		ByName(city string) (*OpenWeatherResponse, error)
	}

	openWeather struct {
		token string
	}

	OpenWeatherResponse struct {
		Coord struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Base string `json:"base"`
		Main struct {
			Temp     float64 `json:"temp"`
			Pressure float64 `json:"pressure"`
			Humidity int     `json:"humidity"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
		} `json:"main"`
		Visibility int `json:"visibility"`
		Wind       struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Dt  int `json:"dt"`
		Sys struct {
			Type    int     `json:"type"`
			ID      int     `json:"id"`
			Message float64 `json:"message"`
			Country string  `json:"country"`
			Sunrise int     `json:"sunrise"`
			Sunset  int     `json:"sunset"`
		} `json:"sys"`
		ID   int    `json:"id"`
		Name string `json:"name"`
		Cod  int    `json:"cod"`
	}
)

func (r *OpenWeatherResponse) DescriptionTotal() string {
	total := []string{}
	for _, v := range r.Weather {
		total = append(total, v.Description)
	}

	return strings.Join(total, " ")
}

func NewWeather(token string) Weather {
	return &openWeather{
		token: token,
	}
}

//ByName returns weather by City name.
func (w *openWeather) ByName(city string) (*OpenWeatherResponse, error) {
	body, err := w.makeRequest(city)
	if err != nil {
		return nil, err
	}

	return w.transformStringToResponse(body)
}

func (w *openWeather) makeRequest(city string) (body string, err error) {
	url := fmt.Sprintf(openWeatherURL, city, w.token)
	logrus.Infof("Making Request: %s", url)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	logrus.Infof("OpenWeatherResponse body: %s", body)
	return body, nil
}

func (w *openWeather) transformStringToResponse(body string) (resp *OpenWeatherResponse, err error) {
	err = json.Unmarshal([]byte(body), &resp)
	return
}
