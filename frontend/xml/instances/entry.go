package instances

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Entry struct {
	Node
	Title     string
	Content   string
	Id        string
	Comments  string
	Enclosure string
	Source    string

	Dates      []Date
	Categories []string
	Medias     []Media
	Links      []Link
	People     []Person
}

func (t *Entry) SetChardata(d xml.CharData) { /* skip */ }

func CreateEntry() *Entry {
	v := Entry{}
	v.NodeTag = "Entry"
	return &v
}

func (t *Entry) Append(o NodeI) {
	switch c := o.(type) {
	case *Content:
		c.Simplify()
		t.Content = c.Value
		t.Medias = append(t.Medias, c.Medias...)
	case *Title:
		t.Title = (*c).Value
	case *Id:
		t.Id = (*c).Value
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
	default:
		utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
	}
}
