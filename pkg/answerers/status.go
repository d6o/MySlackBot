package answerers

import (
	"fmt"

	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/fsouza/go-dockerclient"
)

const (
	endpoint            = "unix:///var/run/docker.sock"
	statusMessageFormat = "%s - %s %s"
)

func Status(message slack.Message) (messageList []slack.Message, err error) {
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		return
	}
	for _, container := range containers {
		message.Text = fmt.Sprintf(statusMessageFormat, container.Names, container.Status, container.State)
		messageList = append(messageList, message)
	}
	return
}
