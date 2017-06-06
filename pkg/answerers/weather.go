package answerers

import (
	"fmt"

	"github.com/disiqueira/MySlackBot/pkg/answerers/weather"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

const (
	weatherAnswerFormat = "%s, %s - Current: %s %-2.0fC, Humidity: %d%% High: %-2.0fC, Low: %-2.0fC"
)

var (
	weatherUserPrefs map[string]string
)

func init() {
	weatherUserPrefs = make(map[string]string)
}

//Weather TODO
func Weather(message slack.Message, weather *weather.OpenWeather) (answer slack.Message, err error) {
	answer = message
	city := bestCity(message)
	resp, err := weather.ByName(city)
	if err != nil {
		return answer, err
	}
	weatherUserPrefs[message.User] = city
	answer.Text = fmt.Sprintf(weatherAnswerFormat, resp.Name, resp.Sys.Country, resp.DescriptionTotal(), resp.Main.Temp, resp.Main.Humidity, resp.Main.TempMax, resp.Main.TempMin)
	return
}

func bestCity(message slack.Message) string {
	if len(message.Text) > 4 {
		return message.Text
	}
	_, prs := weatherUserPrefs[message.User]
	if prs {
		return weatherUserPrefs[message.User]
	}
	return "São Paulo"
}
