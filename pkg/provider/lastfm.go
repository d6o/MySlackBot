package provider

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

const (
	lastFMURL = "https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json"
)

type (
	LastFM interface {
		ByUser(string) (*LastFMResponse, error)
	}

	LastFMResponse struct {
		Recenttracks struct {
			Track []Track `json:"track"`
			Attr  struct {
				User       string `json:"user"`
				Page       string `json:"page"`
				PerPage    string `json:"perPage"`
				TotalPages string `json:"totalPages"`
				Total      string `json:"total"`
			} `json:"@attr"`
		} `json:"recenttracks"`
	}

	Track struct {
		Artist struct {
			Text string `json:"#text"`
			Mbid string `json:"mbid"`
		} `json:"artist"`
		Name       string `json:"name"`
		Streamable string `json:"streamable"`
		Mbid       string `json:"mbid"`
		Album      struct {
			Text string `json:"#text"`
			Mbid string `json:"mbid"`
		} `json:"album"`
		URL   string `json:"url"`
		Image []struct {
			Text string `json:"#text"`
			Size string `json:"size"`
		} `json:"image"`
		Attr struct {
			Nowplaying string `json:"nowplaying"`
		} `json:"@attr,omitempty"`
		Date struct {
			Uts  string `json:"uts"`
			Text string `json:"#text"`
		} `json:"date,omitempty"`
	}
)

//LastTrack TODO
func (r *LastFMResponse) LastTrack() (Track, error) {
	if len(r.Recenttracks.Track) == 0 {
		return Track{}, errors.New("user not found")
	}
	return r.Recenttracks.Track[len(r.Recenttracks.Track)-1], nil
}

//RecentTracks TODO
type RecentTracks struct {
	token string
}

//New TODO
func NewLastFM(token string) LastFM {
	return &RecentTracks{
		token: token,
	}
}

//ByUser returns lastfm recent music of a User.
func (w *RecentTracks) ByUser(user string) (*LastFMResponse, error) {
	body, err := w.makeRequest(user)
	if err != nil {
		return nil, err
	}

	return w.transformStringToResponse(body)
}

func (w *RecentTracks) makeRequest(city string) (body string, err error) {
	url := fmt.Sprintf(lastFMURL, city, w.token)
	logrus.Infof("Making Request: %s", url)
	_, body, errs := gorequest.New().Get(url).End()
	if errs != nil {
		return "", errs[0]
	}
	logrus.Infof("OpenWeatherResponse body: %s", body)
	return body, nil
}

func (w *RecentTracks) transformStringToResponse(body string) (resp *LastFMResponse, err error) {
	err = json.Unmarshal([]byte(body), &resp)
	return
}
