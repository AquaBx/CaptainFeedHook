package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func elementOpened(pile *utils.Stack[NodeI], tag xml.StartElement) {
	var newEl NodeI
	if tag.Name.Local == "rss" {
		return
	} else if tag.Name.Local == "channel" || tag.Name.Local == "feed" {
		newEl = CreateChannel()
	} else if tag.Name.Local == "entry" || tag.Name.Local == "item" {
		newEl = CreateEntry()
	} else if tag.Name.Local == "subtitle" {
		newEl = CreateSubtitle()
	} else if tag.Name.Local == "link" {
		newEl = CreateLink()
	} else if tag.Name.Local == "category" {
		newEl = CreateCategory()
	} else if tag.Name.Local == "title" {
		newEl = CreateTitle()
	} else if tag.Name.Local == "name" {
		newEl = CreateName()
	} else if (tag.Name.Space == "http://www.w3.org/2005/Atom" && tag.Name.Local == "content") || (tag.Name.Space == "" && tag.Name.Local == "description") || (tag.Name.Space == "http://www.w3.org/2005/Atom" && tag.Name.Local == "description") || (tag.Name.Local == "description" && tag.Name.Space == "http://search.yahoo.com/mrss/") {
		newEl = CreateContent()
	} else if tag.Name.Local == "encoded" && tag.Name.Space == "http://purl.org/rss/1.0/modules/content/" {
		a := CreateContent()
		a.Type = "html"
		newEl = a
	} else if tag.Name.Local == "id" || tag.Name.Local == "guid" {
		newEl = CreateId()
	} else if tag.Name.Local == "published" || tag.Name.Local == "pubDate" {
		newEl = CreatePublished()
	} else if tag.Name.Local == "updated" || tag.Name.Local == "lastBuildDate" {
		newEl = CreateUpdated()
	} else if tag.Name.Local == "author" || tag.Name.Local == "creator" {
		newEl = CreateAuthor()
	} else if tag.Name.Local == "thumbnail" {
		newEl = CreateThumbnail()
	} else if tag.Name.Local == "icon" {
		newEl = CreateIcon()
	} else if tag.Name.Local == "copyright" {
		newEl = CreateCopyright()
	} else if tag.Name.Local == "logo" {
		newEl = CreateLogo()
	} else if tag.Name.Space == "http://search.yahoo.com/mrss/" && tag.Name.Local == "group" {
		return
	} else if tag.Name.Space == "" && tag.Name.Local == "generator" {
		newEl = CreateGenerator()
	} else if tag.Name.Space == "http://purl.org/rss/1.0/modules/syndication/" && tag.Name.Local == "updatePeriod" {
		newEl = CreateUpdatePeriod()
	} else if tag.Name.Space == "http://purl.org/rss/1.0/modules/syndication/" && tag.Name.Local == "updateFrequency" {
		newEl = CreateUpdateFrequency()
	} else if tag.Name.Space == "http://search.yahoo.com/mrss/" && tag.Name.Local == "content" {
		newEl = CreateMedia()
	} else if tag.Name.Local == "image" || tag.Name.Local == "enclosure" {
		newEl = CreateImage()
	} else if tag.Name.Local == "uri" || tag.Name.Local == "url" {
		newEl = CreateUri()
	} else if tag.Name.Local == "language" {
		newEl = CreateLanguage()
	} else if tag.Name.Local == "width" {
		newEl = CreateWidth()
	} else if tag.Name.Local == "height" {
		newEl = CreateHeight()
	} else {
		newEl = CreateNode(tag.Name.Local + " " + tag.Name.Space)
	}

	for _, attr := range tag.Attr {
		newEl.SetAttr(attr)
	}

	pile.Push(newEl)
}

func elementChardata(pile *utils.Stack[NodeI], chardata xml.CharData) {
	v := pile.Pop()
	v.SetChardata(chardata)
	pile.Push(v)
}

func elementClosed(pile *utils.Stack[NodeI]) {
	value := pile.Pop()
	parent := pile.Pop()
	parent.Append(value)
	pile.Push(parent)
}

func FetchRSS(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	if err != nil {
		utils.Log("error", err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.Log("error", err.Error())
	}

	return body, nil
}

func Visitor(body []byte) RSS {
	decoder := xml.NewDecoder(strings.NewReader(string(body[:])))

	pile := utils.Stack[NodeI]{}
	pile.Push(CreateRSS())

	shouldBreak := false

	for {

		if shouldBreak {
			break
		}

		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		switch tag := token.(type) {
		case xml.StartElement:
			if tag.Name.Local == "script" {
				shouldBreak = true
				break
			}
			elementOpened(&pile, tag)
		case xml.EndElement:
			if tag.Name.Local == "rss" {
				shouldBreak = true
			} else if tag.Name.Space == "http://search.yahoo.com/mrss/" && tag.Name.Local == "group" {
			} else {
				elementClosed(&pile)
			}
		case xml.CharData:
			elementChardata(&pile, tag)
		case xml.Comment:
			utils.Log("debug", fmt.Sprintf("Comment: %+v", tag))
		case xml.ProcInst:
			utils.Log("debug", fmt.Sprintf("ProcInst: %s", tag))
		case xml.Directive:
			utils.Log("debug", fmt.Sprintf("Directive: %+v", tag))
		default:
			utils.Log("debug", fmt.Sprintf("Comment: %+v", tag))

		}
	}

	return *(pile.Pop().(*RSS))
}
