package slack

import (
	"sync/atomic"

	"strings"

	"github.com/disiqueira/MySlackBot/pkg/slack/rtm"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

const (
	originURL         = "https://api.slack.com/"
	textType          = "message"
	counterInitial    = 1
	counterIncrement  = 1
	webSocketProtocol = ""
	spamMessage       = "Pssst! I didn’t unfurl"
)

type (
	Agent interface {
		Message() (Message, error)
		SendMessage(Message)
	}

	agent struct {
		ws      *websocket.Conn
		counter uint64
	}

	Message struct {
		ID      uint64 `json:"id"`
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"text"`
		User    string `json:"user"`
		Ts      string `json:"ts"`
	}
)

func New(rtm *rtm.Response) (Agent, error) {
	a := &agent{
		counter: counterInitial,
	}
	return a, a.connect(rtm.URL)
}

func (a *agent) connect(rtm string) error {
	var err error
	a.ws, err = websocket.Dial(rtm, webSocketProtocol, originURL)
	return err
}

func (a *agent) message() (Message, error) {
	var m Message
	if err := websocket.JSON.Receive(a.ws, &m); err != nil {
		return m, err
	}

	logrus.Infof("Message: %v", m)
	return m, nil
}

func (a *agent) Message() (Message, error) {
	for {
		m, err := a.message()
		if err != nil {
			return m, err
		}

		if strings.Contains(m.Text, spamMessage) {
			continue
		}

		if m.Type == textType {
			m.Text = strings.ToLower(m.Text)
			return m, nil
		}
	}
}

//SendMessage sends a Message to a channel.
func (a *agent) SendMessage(m Message) {
	m.ID = atomic.AddUint64(&a.counter, counterIncrement)
	go websocket.JSON.Send(a.ws, m)
}
