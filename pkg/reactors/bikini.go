package reactors

import (
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	Bikini interface {
		listener.Reactor
	}

	bikini struct {
		prefix           string
		instagram        provider.Instagram
		imageRecognition provider.ImageRecognition
	}
)

const (
	numPhotos       = 100
	concept         = "bikini"
	minConceptValue = 0.9
)

func NewBikini(instagram provider.Instagram, imageRecognition provider.ImageRecognition, prefix string) Bikini {
	return &bikini{
		prefix:           prefix,
		instagram:        instagram,
		imageRecognition: imageRecognition,
	}
}

func (b *bikini) Usage() string {
	return b.prefix + " {user}"
}

func (b *bikini) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, b.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, b.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	photos, err := b.filterPhotos(b.userFromMessageText(text), concept, minConceptValue)
	if err != nil {
		return err
	}

	for _, photo := range photos {
		answer.Text = photo
		agent.SendMessage(answer)
	}
	return nil
}

func (b *bikini) filterPhotos(user, conceptFilter string, valueFilter float64) ([]string, error) {
	photos, err := b.instagram.LastPhotos(user, numPhotos)
	if err != nil {
		return nil, err
	}

	concepts, err := b.imageRecognition.Analyze(photos)
	if err != nil {
		return nil, err
	}

	final := []string{}
	for url, concept := range concepts {
		val, ok := concept[conceptFilter]
		if val > valueFilter && ok {
			final = append(final, url)
		}
	}
	return final, nil
}

func (b *bikini) userFromMessageText(text string) string {
	if len(text) == 0 {
		return userDefault
	}

	return text
}
