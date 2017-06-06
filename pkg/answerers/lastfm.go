package answerers

import (
	"fmt"

	"github.com/disiqueira/MySlackBot/pkg/answerers/lastfm"
	"github.com/disiqueira/MySlackBot/pkg/slack"
)

const (
	lastfmAnswerFormat = "%s is listening to \"%s\" by %s from the album %s."
)

var (
	lastfmUserPrefs map[string]string
)

func init() {
	lastfmUserPrefs = make(map[string]string)
}

//LastFM TODO
func LastFM(message slack.Message, lastfm *lastfm.RecentTracks) (answer slack.Message, err error) {
	answer = message
	user := bestUser(message)
	resp, err := lastfm.ByUser(user)
	if err != nil {
		return answer, err
	}
	lastfmUserPrefs[message.User] = user

	lastTrack, err := resp.LastTrack()
	if err != nil {
		answer.Text = "User not found."
		return answer, nil
	}

	answer.Text = fmt.Sprintf(lastfmAnswerFormat, resp.Recenttracks.Attr.User, lastTrack.Name, lastTrack.Artist.Text, lastTrack.Album.Text)
	return
}

func bestUser(message slack.Message) string {
	if len(message.Text) > 4 {
		return message.Text
	}
	_, prs := lastfmUserPrefs[message.User]
	if prs {

		return lastfmUserPrefs[message.User]
	}
	return "maef_5"
}
