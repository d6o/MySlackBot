package main

import (
	"fmt"
	"os"

	"github.com/disiqueira/MySlackBot/pkg/config"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/reactors"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/disiqueira/MySlackBot/pkg/slack/rtm"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting MySlackBot")

	logrus.Info("Loading configs")
	var configs config.Specs
	if err := envconfig.Process("msb", &configs); err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Infof("Configs: %v", configs)

	fmt.Println("MySlackBot running!")

	logrus.Info("Starting Slack")

	realTime, err := rtm.New(configs.SlackToken())
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	agent, err := slack.New(realTime)
	if err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}

	weatherProvider := provider.NewWeather(configs.OpenWeatherToken())
	lastFMProvider := provider.NewLastFM(configs.LastFMToken())
	wolframProvider := provider.NewWolfram(configs.WolframToken())
	pokemonProvider := provider.NewPokemon()
	imageProvider := provider.NewImageRecognition(configs.ClarifaiToken())

	consumer := listener.NewConsumer(agent)
	consumer.RegisterReactor(reactors.NewList("list"))
	consumer.RegisterReactor(reactors.NewChoose("choose"))
	consumer.RegisterReactor(reactors.NewWeather(weatherProvider, "weather", "Bauru"))
	consumer.RegisterReactor(reactors.NewLastFM(lastFMProvider, "lastfm", "maef_5"))
	consumer.RegisterReactor(reactors.NewWolfram(wolframProvider, "wolfram"))
	consumer.RegisterReactor(reactors.NewPokemon(pokemonProvider, "pokemon"))
	consumer.RegisterReactor(reactors.NewImageRecognition(imageProvider, "recog"))
	consumer.RegisterReactor(reactors.NewRandom("random"))

	if err := consumer.Listen(); err != nil {
		logrus.Fatal(err.Error())
		os.Exit(1)
	}
}
