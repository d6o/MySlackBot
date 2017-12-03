package reactors

import (
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"fmt"
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

const (
	multivac = "INSUFFICIENT DATA FOR A MEANINGFUL ANSWER"
)

func NewWolfram(provider provider.Wolfram, prefix string) Wolfram {
	return &wolfram{
		prefix:   prefix,
		provider: provider,
	}
}

func (w *wolfram) Usage() string {
	return w.prefix + " {question} - Answer any question in natural language"
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

	var plainTextList []string
	for _, pod := range wolframAnswer.Queryresult.Pods {
		plainTextList = filterAndAddToList(plainTextList,fmt.Sprintf("*%s:*", pod.Title))
		for _, subPod := range pod.Subpods {
			plainTextList = filterAndAddToList(plainTextList, fmt.Sprintf("*%s:*", subPod.Title))
			plainTextList = filterAndAddToList(plainTextList, subPod.Plaintext)
		}
	}

	for _, mean := range wolframAnswer.Queryresult.Didyoumeans {
		plainTextList = append(plainTextList,fmt.Sprintf("You should try: %s (%s)", mean.Val, mean.Score))
	}

	if len(plainTextList) == 0 {
		answer.Text = multivac

		agent.SendMessage(answer)
		return nil
	}

	answer.Text = strings.Join(plainTextList, "\n")

	agent.SendMessage(answer)
	return nil
}

func filterAndAddToList(list []string, item string) []string {
	item = strings.Replace(item, "|", ":", -1)
	item = strings.TrimSpace(item)
	if item != "" && len(item) > 3 {
		return append(list, item)
	}
	return list
}
