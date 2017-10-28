package reactors

import (
	"math/rand"
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	Choose interface {
		listener.Reactor
	}

	choose struct {
		prefix string
	}
)

func NewChoose(prefix string) Choose {
	return &choose{
		prefix: prefix,
	}
}

func (c *choose) Usage() string {
	return c.prefix + " {param1}, {param2} [, {param3} ...]"
}

func (c *choose) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, c.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, c.prefix, "", 1)
	text = strings.Trim(text, " ")

	parts := strings.Split(message.Text, ",")
	answer := message
	answer.Text = parts[rand.Intn(len(parts))]

	agent.SendMessage(answer)
	return nil
}
