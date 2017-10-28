package reactors

import (
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"strings"
)

type (
	Wolfram interface {
		listener.Reactor
	}

	wolfram struct {
		prefix   string
		provider provider.Wolfram
	}
)

func NewWolfram(provider provider.Wolfram, prefix string) Wolfram {
	return &wolfram{
		prefix:   prefix,
		provider: provider,
	}
}

func (w *wolfram) Usage(agent slack.Agent, message slack.Message) {
	answer := message
	answer.Text = w.prefix + " {question}"
	agent.SendMessage(answer)
}

func (w *wolfram) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, w.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, w.prefix, "", 1)
	text = strings.Trim(text, " ")

	wolframAnswer, err := w.provider.Ask(text)
	if err != nil {
		return err
	}
	answer := message
	answer.Text = wolframAnswer

	agent.SendMessage(answer)
	return nil
}
