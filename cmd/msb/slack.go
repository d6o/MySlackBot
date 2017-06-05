package main

import (
	"github.com/disiqueira/MySlackBot/pkg/answerers"
	"github.com/disiqueira/MySlackBot/pkg/answerers/weather"
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
