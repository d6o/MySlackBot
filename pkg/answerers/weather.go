package answerers

import (
	"fmt"

	"github.com/disiqueira/MySlackBot/pkg/answerers/weather"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

const (
	answerFormat = "%s, %s - Current: %s %-2.0fC, Humidity: %d%% High: %-2.0fC, Low: %-2.0fC"
)

//Weather TODO
func Weather(message slack.Message, weather *weather.OpenWeather) (answer slack.Message, err error) {
	answer = message
	resp, err := weather.ByName(message.Text)
	if err != nil {
		return answer, err
	}

	answer.Text = fmt.Sprintf(answerFormat, resp.Name, resp.Sys.Country, resp.DescriptionTotal(), resp.Main.Temp, resp.Main.Humidity, resp.Main.TempMax, resp.Main.TempMin)
	return
}
