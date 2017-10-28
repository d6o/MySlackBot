package config

type (
	Specification interface {
		SlackToken() string
		OpenWeatherToken() string
		LastFMToken() string
		WolframToken() string
		ClarifaiToken() string
		InstagramUsername() string
		InstagramPassword() string
	}

	Specs struct {
		Slack         string `envconfig:"slack_token" required:"true"`
		OpenWeather   string `envconfig:"openweather_token" required:"true"`
		LastFM        string `envconfig:"lastfm_token" required:"true"`
		Wolfram       string `envconfig:"wolfram_token" required:"true"`
		Clarifai      string `envconfig:"clarifai_token" required:"true"`
		InstagramUser string `envconfig:"instagram_username" required:"true"`
		InstagramPass string `envconfig:"instagram_password" required:"true"`
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

func (s *Specs) ClarifaiToken() string {
	return s.Clarifai
}

func (s *Specs) InstagramUsername() string {
	return s.InstagramUser
}

func (s *Specs) InstagramPassword() string {
	return s.InstagramPass
}
