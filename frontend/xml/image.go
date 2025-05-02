package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Media struct {
	Node
	Uri   string
	Title string
	Type  string
}

type Thumbnail = Media
type Icon = Media
type Logo = Media
type Image = Media

func CreateMedia(t ...string) *Media {
	v := Media{}
	v.NodeTag = "Media"
	if len(t) > 0 {
		v.NodeTag = t[0]
	}
	return &v
}

func CreateImage() *Image {
	return CreateMedia("Image")
}

func CreateIcon() *Icon {
	return CreateMedia("Icon")
}

func CreateLogo() *Logo {
	return CreateMedia("Logo")
}

func CreateThumbnail() *Thumbnail {
	return CreateMedia("Thumbnail")
}

func (t *Media) SetAttr(d xml.Attr) {
	if d.Name.Local == "uri" && d.Name.Space == "" {
	} else if d.Name.Local == "length" {
	} else if d.Name.Local == "url" && d.Name.Space == "" {
		t.Uri = d.Value
	} else if d.Name.Local == "type" && d.Name.Space == "" {
		t.Type = d.Value
	} else if d.Name.Local == "width" {
	} else if d.Name.Local == "height" {
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))

	}
}

func (t *Media) SetChardata(d xml.CharData) {
	if t.Uri == "" {
		t.Uri = string(d)
	}
}

func (t *Media) Append(o NodeI) {
	switch c := o.(type) {
	case *Uri:
		t.Uri = (*c).Value
	case *Title:
		t.Title = (*c).Value
	case *Link:
		if t.Uri == "" {
			t.Uri = (*c).Href
		}

	case *Width:
	case *Height:
	default:
		utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
	}
}
