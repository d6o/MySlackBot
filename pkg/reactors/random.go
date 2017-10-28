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
	Random interface {
		listener.Reactor
	}

	random struct {
		prefix string
	}
)

const (
	maxValueDefault = 100
)

func NewRandom(prefix string) Random {
	return &random{
		prefix: prefix,
	}
}

func (r *random) Usage() string {
	return r.prefix + " {maxValue}"
}

func (r *random) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, r.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, r.prefix, "", 1)
	text = strings.Trim(text, " ")

	answer := message
	answer.Text = r.randomValue(text)

	agent.SendMessage(answer)
	return nil
}

func (r *random) randomValue(text string) string {
	return fmt.Sprint(rand.Intn(r.maxValue(text)))
}

func (r *random) maxValue(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil || val == 0 {
		return maxValueDefault
	}
	return val
}
