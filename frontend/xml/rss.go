package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type RSS struct {
	Node
	Channels []Channel
}

func (t *RSS) SetChardata(d xml.CharData) { /* skip */ }

func CreateRSS() *RSS {
	v := RSS{}
	v.NodeTag = "RSS"
	return &v
}

func (t *RSS) Append(o NodeI) {
	switch o.GetType() {
	case "Channel":
		t.Channels = append(t.Channels, *(o.(*Channel)))
	default:
		utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
	}
}
