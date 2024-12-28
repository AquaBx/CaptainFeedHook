package instances

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Category struct {
	Node
	Value string
	Label string
	Term  string
}

func CreateCategory() *Category {
	v := Category{}
	v.NodeTag = "Category"
	return &v
}

func (t *Category) SetChardata(d xml.CharData) {
	t.Value = string(d)
}

func (t *Category) SetAttr(d xml.Attr) {
	if d.Name.Local == "label" && d.Name.Space == "" {
		t.Label = d.Value
	} else if d.Name.Local == "term" && d.Name.Space == "" {
		t.Term = d.Value
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
	}
}
