package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordImage struct {
	Url string `json:"url"`
}

type DiscordEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       int    `json:"color"`
	Timestamp   string `json:"timestamp"`
	Url         string `json:"url"`

	Image DiscordImage `json:"image"`
}

type DiscordWebhook struct {
	Username   string `json:"username"`
	ThreadName string `json:"thread_name"`
	AvatarUrl  string `json:"avatar_url"`

	Embeds []DiscordEmbed `json:"embeds"`
}

func (b DiscordWebhook) Send(uri string) error {
	if marshal, err := json.Marshal(b); err != nil {
		return err
	} else {
		resp, err := http.Post(uri, "application/json", bytes.NewBuffer(marshal))
		if err != nil {
			return err
		}
		if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
			return fmt.Errorf("%s", resp.Status)
		}
	}
	return nil
}
