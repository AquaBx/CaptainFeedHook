package instances

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
	"strconv"
)

type Node struct {
	NodeTag string
}

type NodeI interface {
	GetType() string
	SetChardata(d xml.CharData)
	Append(o NodeI)
	SetAttr(a xml.Attr)
	Simplify()
}

func (t *Node) GetType() string {
	return t.NodeTag
}

func (t *Node) Simplify() {}

func (t *Node) Append(o NodeI) {
	utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
}

func (t *Node) SetChardata(d xml.CharData) {
	utils.Log("debug", fmt.Sprintf("%s.SetChardata does nothing", t.GetType()))
}

func (t *Node) SetAttr(d xml.Attr) {
	utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
}

func CreateNode(t string) *Node {
	v := Node{}
	v.NodeTag = "Node(" + t + ")"
	return &v
}

func parseInt(d string) int {
	v, err := strconv.Atoi(d)

	if err != nil {
		panic(err)
	}
	return v
}
