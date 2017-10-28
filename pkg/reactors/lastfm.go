package reactors

import (
	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"strings"
)

type (
	LastFM interface {
		listener.Reactor
	}

	lastFM struct {
		prefix   string
		provider provider.LastFM
		fallback string
	}
)

const (
	lastFMAnswerFormat = "%s is listening to \"%s\" by %s from the album %s."
)

func NewLastFM(provider provider.LastFM, prefix, fallback string) LastFM {
	return &lastFM{
		prefix:   prefix,
		provider: provider,
		fallback: fallback,
	}
}

func (l *lastFM) Usage(agent slack.Agent, message slack.Message) {
	answer := message
	answer.Text = l.prefix + " {user}"
	agent.SendMessage(answer)
}

func (l *lastFM) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, l.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, l.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	user := l.bestUser(text)
	resp, err := l.provider.ByUser(user)
	if err != nil {
		return err
	}

	lastTrack, err := resp.LastTrack()
	if err != nil {
		answer.Text = "User not found."
		agent.SendMessage(answer)
		return nil
	}

	answer.Text = fmt.Sprintf(lastFMAnswerFormat, resp.Recenttracks.Attr.User, lastTrack.Name, lastTrack.Artist.Text, lastTrack.Album.Text)
	agent.SendMessage(answer)
	return nil
}

func (l *lastFM) bestCity(message string) string {
	if len(message) > 4 {
		return message
	}

	return l.fallback
}

func (l *lastFM) bestUser(message string) string {
	if len(message) > 4 {
		return message
	}
	return l.fallback
}
