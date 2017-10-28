package reactors

import (
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
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

func (w *wolfram) Usage() string {
	return w.prefix + " {question}"
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
