package main

import (
	"CaptainFeedHook/backend"
	"CaptainFeedHook/frontend/xml"
	"CaptainFeedHook/frontend/xml/instances"
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

// An IntHeap is a min-heap of ints.
type Priority []PriorityItem

type PriorityItem struct {
	Id       string
	NextCall int64
	LastCall int64
}

func (h Priority) Len() int           { return len(h) }
func (h Priority) Less(i, j int) bool { return h[i].NextCall < h[j].NextCall }
func (h Priority) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Priority) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PriorityItem))
}

func (h *Priority) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func SearchImage(medias []instances.Media) string {
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

func SearchDate(dates []instances.Date) int64 {
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
	save := Priority{}
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
			heap.Push(&save, PriorityItem{Id: id, NextCall: nowinit, LastCall: nowinit})
		}
	}

	b, _ := json.Marshal(&save)
	saveF.Write(b)

	for {
		actual := heap.Pop(&save).(PriorityItem)
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
