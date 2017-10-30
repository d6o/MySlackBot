package listener

import (
	"errors"

	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	Consumer interface {
		Listen() error
		RegisterReactor(Reactor)
	}

	consumer struct {
		slack    slack.Agent
		reactors []Reactor
	}

	Reactor interface {
		Execute(slack.Agent, slack.Message) error
		Usage() string
	}
)

func NewConsumer(slack slack.Agent) Consumer {
	return &consumer{
		slack: slack,
	}
}

func (o *consumer) Listen() error {
	for {
		m, err := o.slack.Message()
		if err != nil {
			return err
		}
		if err := o.executeReactors(m); err != nil {
			answer := m
			answer.Text = fmt.Sprintf("ERR: %s", err.Error())
			o.slack.SendMessage(answer)
		}
		o.verifyList(m)
	}

	return errors.New("stopped listen for new messages")
}

func (o *consumer) executeReactors(m slack.Message) error {
	for _, reactor := range o.reactors {
		if err := reactor.Execute(o.slack, m); err != nil {
			return err
		}
	}
	return nil
}

func (o *consumer) RegisterReactor(r Reactor) {
	o.reactors = append(o.reactors, r)
}

func (o *consumer) verifyList(m slack.Message) {
	if m.Text == "list" {
		for _, reactor := range o.reactors {
			answer := m
			answer.Text = reactor.Usage()
			o.slack.SendMessage(answer)
		}
	}
}
