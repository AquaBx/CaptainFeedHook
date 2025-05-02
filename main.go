package main

import (
	"CaptainFeedHook/backend"
	"CaptainFeedHook/frontend/xml"
	"CaptainFeedHook/utils"
	"container/heap"
	"encoding/json"
	"fmt"
	"time"
)

type Configs map[string]Config

type Config struct {
	Webhook   string
	Rss       string
	Color     int
	Interval  int64
	AvatarUrl string
	MaxLength int
}

func SearchImage(medias []xml.Media) string {
	uri := ""
	priority := 0

	for _, img := range medias {
		nodePriority := 0
		switch img.NodeTag {
		case "Image":
			nodePriority = 4
		case "Thumbnail":
			nodePriority = 3
		case "Logo":
			nodePriority = 2
		case "Icon":
			nodePriority = 1
		}
		if nodePriority > priority {
			priority = nodePriority
			uri = img.Uri
		}
	}

	return uri
}

func SearchDate(dates []xml.Date) int64 {
	dateV := int64(0)
	priority := 0

	for _, date := range dates {
		nodePriority := 0
		switch date.NodeTag {
		case "Published":
			nodePriority = 2
		case "Updated":
			nodePriority = 1
		}
		if nodePriority > priority {
			priority = nodePriority
			dateV = date.Value
		}
	}

	return dateV
}

func main() {
	utils.InitFlags()
	utils.InitLogger()

	configF := utils.FileM{Directory: "config/config.json"}
	saveF := utils.FileM{Directory: "config/save.json"}

	configs := Configs{}
	save := utils.Priority{}
	json.Unmarshal(configF.Read(), &configs)
	json.Unmarshal(saveF.Read(), &save)

	heap.Init(&save)

	// init
	nowinit := time.Now().UTC().Unix()

	for id := range configs {
		inside := false
		for _, itemS := range save {
			if id == itemS.Id {
				inside = true
				break
			}
		}
		if !inside {
			utils.Log("info", fmt.Sprintf("Loading %s", id))
			heap.Push(&save, utils.PriorityItem{Id: id, NextCall: nowinit, LastCall: nowinit})
		}
	}

	b, _ := json.Marshal(&save)
	saveF.Write(b)

	for {
		actual := heap.Pop(&save).(utils.PriorityItem)
		config := configs[actual.Id]
		now := time.Now().UTC().Unix()

		sleepD := time.Duration(actual.NextCall-now) * time.Second

		utils.Log("info", fmt.Sprintf("Sleeping %s", sleepD))

		time.Sleep(sleepD)

		body, err := xml.FetchRSS(config.Rss)

		if err != nil {
			utils.Log("error", err.Error())
			actual.NextCall = time.Now().UTC().Unix() + 5
			heap.Push(&save, actual)
			b, _ := json.Marshal(&save)
			saveF.Write(b)
			continue
		}

		nextLastCall := actual.LastCall

		utils.Log("info", fmt.Sprintf("Checking %s", actual.Id))

		xmlv := xml.Visitor(body)

		for _, Channel := range xmlv.Channels {
			for _, Entry := range Channel.Entries {
				date := SearchDate(Entry.Dates)
				nextLastCall = max(nextLastCall, date)
				if actual.LastCall < date {
					utils.Log("info", fmt.Sprintf("Loading %s", Entry.Title))

					description := Entry.Content
					if len(description) > config.MaxLength {
						description = description[:config.MaxLength]
					}

					titleThread := Entry.Title
					if len(titleThread) > 97 {
						titleThread = titleThread[:97] + "..."
					}

					embed := backend.DiscordEmbed{
						Description: description,
						Image: backend.DiscordImage{
							Url: SearchImage(Entry.Medias),
						},
						Timestamp: time.Unix(date, int64(0)).Format(time.RFC3339),
						Title:     Entry.Title,
						Color:     config.Color,
						Url:       Entry.Links[0].Href,
					}
					Webhook := backend.DiscordWebhook{
						Username:   Channel.Title,
						AvatarUrl:  config.AvatarUrl,
						ThreadName: titleThread,
						Embeds:     []backend.DiscordEmbed{embed},
					}

					err := Webhook.Send(config.Webhook)

					if err != nil {
						utils.Log("debug", fmt.Sprint(Webhook))
						utils.Log("error", err.Error())
					}

					time.Sleep(1 * time.Second)
				}
			}
		}

		actual.NextCall = time.Now().UTC().Unix() + config.Interval
		actual.LastCall = nextLastCall
		heap.Push(&save, actual)
		b, _ := json.Marshal(&save)
		saveF.Write(b)

	}
}
