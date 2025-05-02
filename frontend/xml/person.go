package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Person struct {
	Node
	Name  string
	Uri   string
	Email string
}

type Author = Person
type Contributor = Person

func CreatePerson() *Person {
	v := Person{}
	v.NodeTag = "Person"
	return &v
}

func CreateContributor() *Contributor {
	v := Contributor{}
	v.NodeTag = "Contributor"
	return &v
}
func CreateAuthor() *Author {
	v := Author{}
	v.NodeTag = "Author"
	return &v
}

func (t *Person) Append(o NodeI) {
	switch c := o.(type) {
	case *Uri:
		t.Uri = (*c).Value
	case *Name:
		t.Name = (*c).Value
	case *Email:
		t.Email = (*c).Value
	default:
		utils.Log("debug", fmt.Sprintf("%s.Append does not manage %s", t.GetType(), o.GetType()))
	}
}
func (t *Person) SetChardata(d xml.CharData) {
	if t.Name == "" {
		t.Name = string(d)
	}
}

func (t *Person) SetAttr(d xml.Attr) {
	if d.Name.Local == "name" && d.Name.Space == "" {
		t.Name = d.Value
	} else if d.Name.Local == "url" && d.Name.Space == "" {
		t.Uri = d.Value
	} else if d.Name.Local == "type" {
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
	}
}
