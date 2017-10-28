package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type (
	ImageRecognition interface {
		Analyze(urls []string) (map[string]Concepts, error)
	}

	Concepts map[string]float64

	clarifai struct {
		token string
	}

	clarifaiRequest struct {
		Inputs []clarifaiInput `json:"inputs"`
	}

	clarifaiInput struct {
		Data clarifaiData `json:"data"`
	}

	clarifaiData struct {
		Image clarifaiImage `json:"image"`
	}

	clarifaiImage struct {
		URL string `json:"url"`
	}

	clarifaiResponse struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Outputs []struct {
			ID     string `json:"id"`
			Status struct {
				Code        int    `json:"code"`
				Description string `json:"description"`
			} `json:"status"`
			CreatedAt time.Time `json:"created_at"`
			Model     struct {
				ID         string    `json:"id"`
				Name       string    `json:"name"`
				CreatedAt  time.Time `json:"created_at"`
				AppID      string    `json:"app_id"`
				OutputInfo struct {
					Message string `json:"message"`
					Type    string `json:"type"`
					TypeExt string `json:"type_ext"`
				} `json:"output_info"`
				ModelVersion struct {
					ID        string    `json:"id"`
					CreatedAt time.Time `json:"created_at"`
					Status    struct {
						Code        int    `json:"code"`
						Description string `json:"description"`
					} `json:"status"`
				} `json:"model_version"`
				DisplayName string `json:"display_name"`
			} `json:"model"`
			Input struct {
				ID   string `json:"id"`
				Data struct {
					Image struct {
						URL string `json:"url"`
					} `json:"image"`
				} `json:"data"`
			} `json:"input"`
			Data struct {
				Concepts []struct {
					ID    string  `json:"id"`
					Name  string  `json:"name"`
					Value float64 `json:"value"`
					AppID string  `json:"app_id"`
				} `json:"concepts"`
			} `json:"data"`
		} `json:"outputs"`
	}
)

const (
	clarifaiURL = "https://api.clarifai.com/v2/models/aaa03c23b3724a16a56b629203edc62c/outputs"
	clarifaiOK  = "Ok"
	minValue    = 0.3
)

var (
	ErrNoConceptsFound = errors.New("no concepts found")
)

func NewImageRecognition(token string) ImageRecognition {
	return &clarifai{
		token: token,
	}
}

func (c *clarifai) Analyze(urls []string) (map[string]Concepts, error) {
	body, err := c.makeRequest(c.createRequest(urls))
	if err != nil {
		return nil, err
	}

	response := clarifaiResponse{}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return nil, err
	}

	return c.readConcepts(&response)
}

func (c *clarifai) readConcepts(response *clarifaiResponse) (map[string]Concepts, error) {
	if response.Status.Description != clarifaiOK {
		return nil, ErrNoConceptsFound
	}

	con := map[string]Concepts{}
	for _, outputs := range response.Outputs {
		for _, concept := range outputs.Data.Concepts {

			if _, ok := con[outputs.Input.Data.Image.URL]; !ok {
				con[outputs.Input.Data.Image.URL] = Concepts{}
			}

			if concept.Value > minValue {
				con[outputs.Input.Data.Image.URL][concept.Name] = concept.Value
			}
		}
	}
	return con, nil
}

func (c *clarifai) createRequest(urls []string) *clarifaiRequest {
	r := &clarifaiRequest{}
	for _, url := range urls {
		r.Inputs = append(r.Inputs, clarifaiInput{
			Data: clarifaiData{
				Image: clarifaiImage{
					URL: url,
				},
			},
		})
	}
	return r
}

func (c *clarifai) makeRequest(r *clarifaiRequest) (string, error) {
	_, body, errs := gorequest.New().Post(clarifaiURL).
		Set("Authorization", "Key "+c.token).
		Send(r).
		End()

	if errs != nil {
		return "", errs[0]
	}

	return body, nil
}

func (co Concepts) String() string {
	final := ""

	for key, value := range co {
		final = fmt.Sprintf("%s %s(%6.4f)", final, key, value)
	}

	return final
}
