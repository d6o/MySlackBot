package reactors

import (
	"errors"
	"fmt"
	"github.com/disiqueira/MySlackBot/pkg/listener"
	"github.com/disiqueira/MySlackBot/pkg/provider"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/sirupsen/logrus"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type (
	Pokemon interface {
		listener.Reactor
	}

	pokemon struct {
		prefix   string
		provider provider.Pokemon
	}
)

const (
	maxPokemon       = 500
	pokeAnswerFormat = "%s (%d) %s"
)

func NewPokemon(provider provider.Pokemon, prefix string) Pokemon {
	return &pokemon{
		prefix:   prefix,
		provider: provider,
	}
}

func (w *pokemon) Usage(agent slack.Agent, message slack.Message) {
	answer := message
	answer.Text = w.prefix + " {pokemon}"
	agent.SendMessage(answer)
}

func (w *pokemon) Execute(agent slack.Agent, message slack.Message) error {
	if !strings.HasPrefix(message.Text, w.prefix) {
		return nil
	}
	text := strings.Replace(message.Text, w.prefix, "", 1)
	text = strings.Trim(text, " ")
	answer := message

	poke, err := w.getPokemon(text)
	if err != nil {
		answer.Text = "Pokemon not found."
		agent.SendMessage(answer)
		return nil
	}

	answer.Text = fmt.Sprintf(pokeAnswerFormat, poke.Name, poke.ID, poke.Sprites.FrontDefault)

	agent.SendMessage(answer)
	return nil
}

func (w *pokemon) getPokemon(text string) (poke *provider.PokemonResponse, err error) {
	poke, err = w.findPokemonByNumber(text)
	if err == nil {
		return
	}
	poke, err = w.findPokemonByWords(text)
	if err == nil {
		return
	}
	poke, err = w.provider.Search(strconv.Itoa(rand.Intn(maxPokemon)))
	return
}

func (w *pokemon) findPokemonByNumber(text string) (poke *provider.PokemonResponse, err error) {
	num := w.findNumber(text)
	if num == 0 {
		return poke, errors.New("number not found")
	}
	logrus.Infof("Num: %v", num)

	poke, err = w.provider.Search(strconv.Itoa(num))
	return
}

func (w *pokemon) findNumber(text string) (num int) {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(text, -1)
	logrus.Infof("Numbers: %v", numbers)
	if len(numbers) < 1 {
		return 0
	}
	logrus.Infof("Numbers[0]: %v", numbers[0])
	num, err := strconv.Atoi(numbers[0])
	logrus.Infof("Err: %v", err)
	if err != nil {
		return 0
	}
	return num
}

func (w *pokemon) findPokemonByWords(text string) (poke *provider.PokemonResponse, err error) {
	words := strings.Split(text, " ")
	for _, v := range words {
		poke, err = w.provider.Search(v)
		if err == nil {
			return
		}
	}
	return poke, errors.New("pokemon not found")
}
