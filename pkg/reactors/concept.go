package reactors

import (
	"strings"

	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

type (
	Concept interface {
		listener.Reactor
	}

	concept struct {
		prefix           string
		instagram        provider.Instagram
		imageRecognition provider.ImageRecognition
	}
)

const (
	numPhotos       = 100
	minConceptValue = 0.9
	conceptDefault  = "bikini"
)

func NewBikini(instagram provider.Instagram, imageRecognition provider.ImageRecognition, prefix string) Concept {
	return NewConcept(instagram, imageRecognition, prefix)
}

func NewConcept(instagram provider.Instagram, imageRecognition provider.ImageRecognition, prefix string) Concept {
	return &concept{
		prefix:           prefix,
		instagram:        instagram,
		imageRecognition: imageRecognition,
	}
}

func (b *concept) Usage() string {
	return b.prefix + " {user},{concept} - Search for concepts in pictures from an Instagram account."
}

func (b *concept) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, b.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, b.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	user, concept := b.userAndConceptFromMessageText(text)
	photos, err := b.filterPhotos(user, concept, minConceptValue)
	if err != nil {
		return err
	}

	if len(photos) == 0 {
		answer.Text = fmt.Sprintf("I couldn't find any %s photos in %s account. :(", concept, user)
		agent.SendMessage(answer)
	}

	for _, photo := range photos {
		answer.Text = photo
		agent.SendMessage(answer)
	}
	return nil
}

func (b *concept) filterPhotos(user, conceptFilter string, valueFilter float64) ([]string, error) {
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

func (b *concept) userAndConceptFromMessageText(text string) (string, string) {
	if len(text) == 0 {
		return userDefault, conceptDefault
	}

	split := strings.Split(text, ",")
	if len(split) == 1 {
		return split[0], conceptDefault
	}

	return strings.TrimSpace(split[0]), strings.TrimSpace(split[1])
}
