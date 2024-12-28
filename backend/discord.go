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
	Username    string `json:"username"`
	Thread_name string `json:"thread_name"`
	Avatar_url  string `json:"avatar_url"`

	Embeds []DiscordEmbed `json:"embeds"`
}

func (b DiscordWebhook) Send(uri string) error {
	if json, err := json.Marshal(b); err != nil {
		return err
	} else {
		resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json))
		if err != nil {
			return err
		}
		if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
			return fmt.Errorf("%s", resp.Status)
		}
	}
	return nil
}

/*
{
	"username":"UnrealBlog",
	"thread_name":"Cool",
    "embeds": [
	  {
		"title": "<p>Explore the new updates to the Game Animation Sample Project in UE 5.5</p>",
		"description": "<img src=&quot;https://media.graphassets.com/Qbl8z4ErQFGn0aO3ZFXI&quot; /> The Game Animation Sample Project has been updated for UE 5.5 with over 300 animations, a new Experimental setup that provides better artist control without diminishing quality, setups that are 100% networked and ready for multiplayer, mobile support—and more!",
		"color": 2326507,
        "url": "https://www.unrealengine.com/tech-blog/explore-the-new-updates-to-the-game-animation-sample-project-in-ue-5.5",
		"timestamp": "2024-12-17T23:00:00.000Z"

		"image": {
		  "url": "https://media.graphassets.com/ZgQEnKIIToVf2XeFDKE9"
		},

		"footer": {
		  "text": "Autheur"
		},
	  }
	]
}*/
