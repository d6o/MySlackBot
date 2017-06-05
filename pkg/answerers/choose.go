package answerers

import (
	"math/rand"
	"strings"

	"github.com/disiqueira/MySlackBot/pkg/slack"
)

//Choose TODO
func Choose(message slack.Message) (answer slack.Message) {
	parts := strings.Split(message.Text, ",")
	answer = message
	answer.Text = parts[rand.Intn(len(parts))]
	return
}
