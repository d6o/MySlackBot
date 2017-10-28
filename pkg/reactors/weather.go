package reactors

import (
	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"strings"
)

type (
	Weather interface {
		listener.Reactor
	}

	weather struct {
		prefix   string
		provider provider.Weather
		fallback string
	}
)

const (
	weatherAnswerFormat = "%s, %s - Current: %s %-2.0fC, Humidity: %d%% High: %-2.0fC, Low: %-2.0fC"
)

func NewWeather(provider provider.Weather, prefix, fallback string) Weather {
	return &weather{
		prefix:   prefix,
		provider: provider,
		fallback: fallback,
	}
}

func (w *weather) Usage(agent slack.Agent, message slack.Message) {
	answer := message
	answer.Text = w.prefix + " {city}"
	agent.SendMessage(answer)
}

func (w *weather) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, w.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, w.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	city := w.bestCity(text)
	resp, err := w.provider.ByName(city)
	if err != nil {
		return err
	}

	answer.Text = fmt.Sprintf(weatherAnswerFormat, resp.Name, resp.Sys.Country, resp.DescriptionTotal(), resp.Main.Temp, resp.Main.Humidity, resp.Main.TempMax, resp.Main.TempMin)

	agent.SendMessage(answer)
	return nil
}

func (w *weather) bestCity(message string) string {
	if len(message) > 4 {
		return message
	}

	return w.fallback
}
