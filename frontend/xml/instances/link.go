package instances

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Link struct {
	Node
	Href     string
	Rel      string
	Type     string
	Hreflang string
	Length   string
}

func CreateLink() *Link {
	v := Link{}
	v.NodeTag = "Link"
	return &v
}

func (t *Link) SetChardata(d xml.CharData) {
	if t.Href == "" {
		t.Href = string(d)
	}
}

func (t *Link) SetAttr(d xml.Attr) {
	if d.Name.Local == "rel" && d.Name.Space == "" {
		t.Rel = d.Value
	} else if d.Name.Local == "type" && d.Name.Space == "" {
		t.Type = d.Value
	} else if d.Name.Local == "href" && d.Name.Space == "" {
		t.Href = d.Value
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
	}
}
