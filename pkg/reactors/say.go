package reactors

import (
	"math/rand"
	"strings"

	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"strconv"
)

type (
	Say interface {
		listener.Reactor
	}

	say struct {
		prefix string
	}
)

func NewSay(prefix string) Say {
	return &say{
		prefix: prefix,
	}
}

func (r *say) Usage() string {
	return r.prefix + " {text} - Just repeats the text"
}

func (r *say) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, r.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, r.prefix, "", 1)
	text = strings.Trim(text, " ")

	answer := message
	answer.Text = text

	agent.SendMessage(answer)
	return nil
}
