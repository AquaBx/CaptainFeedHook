package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Channel struct {
	Node
	Title     string
	Copyright string
	Generator string
	Id        string
	Subtitle  string
	Language  string
	Docs      string

	//WebMaster      string
	//ManagingEditor string

	Content         string
	UpdatePeriod    string
	UpdateFrequency string

	Links      []Link
	Dates      []Date
	Categories []string
	Medias     []Media
	People     []Person
	Entries    []Entry
}

func CreateChannel() *Channel {
	v := Channel{}
	v.NodeTag = "Channel"
	return &v
}
func (t *Channel) SetChardata(d xml.CharData) {
	// skip
}
func (t *Channel) Append(o NodeI) {
	switch c := o.(type) {
	case *Entry:
		t.Entries = append(t.Entries, *c)
	case *Language:
		t.Language = (*c).Value
	case *Copyright:
		t.Copyright = (*c).Value
	case *Content:
		c.Simplify()
		t.Content = c.Value
		t.Medias = append(t.Medias, c.Medias...)
	case *Category:
		t.Categories = append(t.Categories, (*c).Value)
	case *Person:
		t.People = append(t.People, *c)
	case *Date:
		t.Dates = append(t.Dates, *c)
	case *Media:
		t.Medias = append(t.Medias, *c)
	case *Link:
		t.Links = append(t.Links, *c)
	case *Title:
		t.Title = (*c).Value
	case *Id:
		t.Id = (*c).Value
	case *Generator:
		t.Generator = (*c).Value
	case *UpdatePeriod:
		t.UpdatePeriod = (*c).Value
	case *UpdateFrequency:
		t.UpdateFrequency = (*c).Value
	case *Subtitle:
		t.Subtitle = (*c).Value

	default:
		utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
	}
}

func (t *Channel) SetAttr(d xml.Attr) {
	if d.Name.Local == "xmlns" || d.Name.Space == "xmlns" {
	} else if d.Name.Local == "lang" {
		t.Language = d.Value
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
	}
}
