package answerers

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/answerers/pokemon"
	"github.com/disiqueira/MySlackBot/pkg/slack"
	"github.com/sirupsen/logrus"
)

const (
	maxPokemon       = 500
	pokeAnswerFormat = "%s (%d) %s"
)

//Pokemon TODO
func Pokemon(message slack.Message) (answer slack.Message) {
	answer = message
	poke, err := getPokemon(message.Text)
	if err != nil {
		answer.Text = "Pokemon not found."
		return answer
	}

	answer.Text = fmt.Sprintf(pokeAnswerFormat, poke.Name, poke.ID, poke.Sprites.FrontDefault)
	return
}

func getPokemon(text string) (poke *pokemon.Response, err error) {
	poke, err = findPokemonByNumber(text)
	if err == nil {
		return
	}
	poke, err = findPokemonByWords(text)
	if err == nil {
		return
	}
	poke, err = pokemon.Search(strconv.Itoa(rand.Intn(maxPokemon)))
	return
}

func findPokemonByNumber(text string) (poke *pokemon.Response, err error) {
	num := findNumber(text)
	if num == 0 {
		return poke, errors.New("Number not found")
	}
	logrus.Infof("Num: %v", num)

	poke, err = pokemon.Search(strconv.Itoa(num))
	return
}

func findNumber(text string) (num int) {
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

func findPokemonByWords(text string) (poke *pokemon.Response, err error) {
	words := strings.Split(text, " ")
	for _, v := range words {
		poke, err = pokemon.Search(v)
		if err == nil {
			return
		}
	}
	return poke, errors.New("Pokemon not found")
}
