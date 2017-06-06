package main

import (
	"github.com/disiqueira/MySlackBot/pkg/answerers"
	"github.com/disiqueira/MySlackBot/pkg/answerers/lastfm"
	"github.com/disiqueira/MySlackBot/pkg/answerers/weather"
	"github.com/disiqueira/MySlackBot/pkg/answerers/wolfram"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/disiqueira/MySlackBot/pkg/slack/rtm"
	"github.com/sirupsen/logrus"
)

func startSlack() (slackAgent *slack.Agent, err error) {
	realTime, err := rtm.New(configs.SlackToken)
	if err != nil {
		return nil, err
	}
	return slack.New(realTime)
}

func chooseCMD() {
	logrus.Info("Starting choose")
	for {
		m, err := slackAgent.PrefixMessage("choose")
		if err != nil {
			logrus.Fatal(err)
		}

		slackAgent.SendMessage(answerers.Choose(m))
	}
}

func weatherCMD() {
	logrus.Info("Starting weather")

	we := weather.New(configs.OpenWeatherToken)

	for {
		m, err := slackAgent.PrefixMessage("we")
		if err != nil {
			logrus.Fatal(err)
		}
		answer, err := answerers.Weather(m, we)
		if err != nil {
			logrus.Fatal(err)
		}

		slackAgent.SendMessage(answer)
	}
}

func lastfmCMD() {
	logrus.Info("Starting LastFM")

	last := lastfm.New(configs.LastFMToken)

	for {
		m, err := slackAgent.PrefixMessage("lastfm")
		if err != nil {
			logrus.Fatal(err)
		}
		answer, err := answerers.LastFM(m, last)
		if err != nil {
			logrus.Fatal(err)
		}

		slackAgent.SendMessage(answer)
	}
}

func wolframCMD() {
	logrus.Info("Starting Wolfram Alpha")

	wolf := wolfram.New(configs.WolframToken)

	for {
		m, err := slackAgent.PrefixMessage("aline")
		if err != nil {
			logrus.Fatal(err)
		}
		answer, err := answerers.Wolfram(m, wolf)
		if err != nil {
			logrus.Fatal(err)
		}

		slackAgent.SendMessage(answer)
	}
}

func pokemonCMD() {
	logrus.Info("Starting Pokemon")
	for {
		m, err := slackAgent.HasWord("pokemon")
		if err != nil {
			logrus.Fatal(err)
		}

		slackAgent.SendMessage(answerers.Pokemon(m))
	}
}
