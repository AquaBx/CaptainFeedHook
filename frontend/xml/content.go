package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Content struct {
	Node
	Value  string
	Type   string
	Medias []Media
}

func CreateContent() *Content {
	v := Content{}
	v.NodeTag = "Content"
	v.Type = "text/plain"
	return &v
}

func (t *Content) SetChardata(d xml.CharData) {
	t.Value = string(d)
}

func (t *Content) SetAttr(d xml.Attr) {
	if d.Name.Local == "type" && d.Name.Space == "" {
		t.Type = d.Value
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))

	}
}

func (t *Content) Simplify() {
	switch t.Type {
	case "text/plain":
	case "html":
		t.HtmlSanitizer()
	default:
		utils.Log("debug", fmt.Sprintf("%s.Content Type not managed %s", t.GetType(), t.Type))
	}
}

func (t *Content) HtmlSanitizer() {
	p := strings.NewReader(t.Value)
	doc, _ := goquery.NewDocumentFromReader(p)

	doc.Find("script").Each(func(i int, el *goquery.Selection) {
		el.Remove()
	})

	doc.Find("img").Each(func(i int, el *goquery.Selection) {
		m := CreateImage()
		m.Uri, _ = el.Attr("src")
		t.Medias = append(t.Medias, *m)
	})

	t.Value = doc.Text()
}
