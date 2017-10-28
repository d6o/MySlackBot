package reactors

import (
	"strings"

	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	ImageRecognition interface {
		listener.Reactor
	}

	imageRecognition struct {
		prefix   string
		provider provider.ImageRecognition
	}
)

func NewImageRecognition(provider provider.ImageRecognition, prefix string) ImageRecognition {
	return &imageRecognition{
		prefix:   prefix,
		provider: provider,
	}
}

func (i *imageRecognition) Usage() string {
	return i.prefix + " {url} returns the main tags of the picture"
}

func (i *imageRecognition) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, i.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, i.prefix, "", 1)
	url := cleanURL(text, []string{" ", ">", "<"})

	fmt.Printf("text: %s \n", url)

	res, err := i.provider.Analyze([]string{url})
	if err != nil {
		if err == provider.ErrNoConceptsFound {
			return nil
		}
		return err
	}

	fmt.Println(res)

	for _, value := range res {
		answer := message
		answer.Text = value.String()
		agent.SendMessage(answer)
	}

	return nil
}

func cleanURL(url string, cutset []string) string {
	final := url
	for _, cut := range cutset {
		final = strings.Trim(final, cut)
	}

	return final
}
