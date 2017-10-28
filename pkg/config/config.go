package config

type (
	Specification interface {
		SlackToken() string
		OpenWeatherToken() string
		LastFMToken() string
		WolframToken() string
	}

	Specs struct {
		Slack       string `envconfig:"slack_token" required:"true"`
		OpenWeather string `envconfig:"openweather_token" required:"true"`
		LastFM      string `envconfig:"lastfm_token" required:"true"`
		Wolfram     string `envconfig:"wolfram_token"`
	}
)

func (s *Specs) SlackToken() string {
	return s.Slack
}
func (s *Specs) OpenWeatherToken() string {
	return s.OpenWeather
}
func (s *Specs) LastFMToken() string {
	return s.LastFM
}
func (s *Specs) WolframToken() string {
	return s.Wolfram
}
