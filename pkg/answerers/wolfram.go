package answerers

import (
	"github.com/disiqueira/MySlackBot/pkg/answerers/wolfram"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

//Wolfram TODO
func Wolfram(message slack.Message, wolfram *wolfram.Short) (answer slack.Message, err error) {
	answer = message

	resp, err := wolfram.Ask(message.Text)
	if err != nil {
		return answer, err
	}
	answer.Text = resp
	return
}
