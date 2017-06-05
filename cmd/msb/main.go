package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/disiqueira/MySlackBot/pkg/config"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	configs    config.Specification
	slackAgent *slack.Agent
)

func init() {
	logrus.Info("Starting MySlackBot")
	logrus.Info("Loading configs")
	err := envconfig.Process("msb", &configs)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Infof("Configs: %v", configs)
	logrus.Info("Starting Slack")

	slackAgent, err = startSlack()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	fmt.Println("MySlackBot running!")
}

func main() {
	cmd := flag.String("cmd", "", "Command to be executed: (choose)")

	flag.Parse()

	logrus.Infof("Command received: %s", *cmd)

	switch *cmd {
	case "choose":
		chooseCMD()
	case "weather":
		weatherCMD()
	}

	logrus.Fatal("Not good...")
	os.Exit(1)
}
