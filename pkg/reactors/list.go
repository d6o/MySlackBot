package reactors

import (
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	List interface {
		listener.Reactor
	}

	list struct {
		prefix string
	}
)

func NewList(prefix string) List {
	return &list{
		prefix: prefix,
	}
}

func (l *list) Usage() string {
	return l.prefix + " - List all commands"
}

func (l *list) Execute(agent slack.Agent, message slack.Message) error {
	return nil
}
