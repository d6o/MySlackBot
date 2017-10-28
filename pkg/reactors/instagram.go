package reactors

import (
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"strconv"
)

type (
	Instagram interface {
		listener.Reactor
	}

	instagram struct {
		prefix   string
		provider provider.Instagram
	}
)

const (
	numPhotosDefault = 5
	userDefault      = "di.siqueira"
)

func NewInstagram(provider provider.Instagram, prefix string) Instagram {
	return &instagram{
		prefix:   prefix,
		provider: provider,
	}
}

func (i *instagram) Usage() string {
	return i.prefix + " {user}[,{num_photos}]"
}

func (i *instagram) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, i.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, i.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	photos, err := i.provider.LastPhotos(i.userAndNumFromMessageText(text))
	if err != nil {
		return err
	}

	for _, photo := range photos {
		answer.Text = photo
		agent.SendMessage(answer)
	}
	return nil
}

func (i *instagram) userAndNumFromMessageText(text string) (string, int) {
	if len(text) == 0 {
		return userDefault, numPhotosDefault
	}

	split := strings.Split(text, ",")
	if len(split) == 1 {
		return split[0], numPhotosDefault
	}

	num, err := strconv.Atoi(split[1])
	if err != nil {
		return split[0], numPhotosDefault
	}

	return split[0], num
}
