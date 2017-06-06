package slack

import (
	"strings"
	"sync/atomic"

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
)

//New TODO
func New(rtm *rtm.Response) (a *Agent, err error) {
	a = &Agent{
		counter: counterInitial,
	}
	err = a.connect(rtm.URL)
	return
}

//Agent TODO
type Agent struct {
	ws      *websocket.Conn
	counter uint64
}

func (a *Agent) connect(rtm string) (err error) {
	a.ws, err = websocket.Dial(rtm, webSocketProtocol, originURL)
	return
}

//Message TODO
type Message struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
	Ts      string `json:"ts"`
}

func (a *Agent) message() (m Message, err error) {
	err = websocket.JSON.Receive(a.ws, &m)
	return
}

func (a *Agent) text() (m Message, err error) {
	m, err = a.message()
	logrus.Info("Message received.")
	logrus.Infof("Message: %v", m)
	if m.Type != textType {
		logrus.Info("Wrong type")
		return a.text()
	}
	return
}

//PrefixMessage returns a Message when it has some prefix string.
func (a *Agent) PrefixMessage(prefix string) (m Message, err error) {
	m, err = a.text()

	logrus.Infof("Text: %s", m.Text)
	logrus.Infof("Prefix: %s", prefix)

	if !strings.HasPrefix(strings.ToLower(m.Text), prefix) {
		logrus.Info("Wrong prefix")
		return a.PrefixMessage(prefix)
	}
	logrus.Info("Removing prefix")
	m.Text = strings.Replace(m.Text, prefix+" ", "", -1)
	m.Text = strings.Trim(m.Text, " ")
	return
}

//SendMessage sends a Message to a channel.
func (a *Agent) SendMessage(m Message) {
	m.ID = atomic.AddUint64(&a.counter, counterIncrement)
	go websocket.JSON.Send(a.ws, m)
}
