package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func SendWebhook(url string) error {
	body := Webhook{
		Content:  "@everyone",
		Username: "Parti Monitor",
		Embeds: []Embed{
			{
				Title:       "Joshua Block is live on Parti!",
				Color:       3136040,
				Description: "https://parti.com/creator/parti/worldoftshirts2001",
				Fields:      nil,
				Footer: struct {
					Text    string `json:"text"`
					IconURL string `json:"icon_url"`
				}{
					Text: "worldoftshirts monitor",
				},
				Thumbnail: struct {
					URL string `json:"url"`
				}{},
			},
		},
		Attachments: nil,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		log.Println("Webhook response body:", string(body))
		return errors.New("failed to send webhook, status code: " + res.Status)
	}

	return nil
}
