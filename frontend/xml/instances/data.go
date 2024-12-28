package instances

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
)

type Data struct {
	Node
	Value string
}

type Number struct {
	Node
	Value int
}

func (t *Data) SetChardata(d xml.CharData) {
	t.Value = string(d)
}

func (t *Number) SetChardata(d xml.CharData) {
	t.Value = parseInt(string(d))
}

func (t *Data) SetAttr(d xml.Attr) {
	if d.Name.Local == "isPermaLink" && d.Name.Space == "" {
		//t = CreateLink()
		//t.Href = d.Value
	} else {
		utils.Log("debug", fmt.Sprintf("%s.SetAttr does not manage %s %s", t.GetType(), d.Name.Space, d.Name.Local))
	}
}

type Name struct{ Data }
type Id struct{ Data }
type Title struct{ Data }
type Copyright struct{ Data }
type UpdateFrequency struct{ Data }
type UpdatePeriod struct{ Data }
type Type struct{ Date }
type Height struct{ Number }
type Width struct{ Number }
type Language struct{ Data }
type Email struct{ Data }
type Subtitle struct{ Data }
type Uri struct{ Data }
type Generator struct{ Data }

func CreateName() *Name {
	v := Name{}
	v.NodeTag = "Name"
	return &v
}

func CreateId() *Id {
	v := Id{}
	v.NodeTag = "Id"
	return &v
}

func CreateGenerator() *Generator {
	v := Generator{}
	v.NodeTag = "Generator"
	return &v
}

func CreateUri() *Uri {
	v := Uri{}
	v.NodeTag = "Uri"
	return &v
}

func CreateTitle() *Title {
	v := Title{}
	v.NodeTag = "Title"
	return &v
}

func CreateSubtitle() *Subtitle {
	v := Subtitle{}
	v.NodeTag = "Subtitle"
	return &v
}

func CreateEmail() *Email {
	v := Email{}
	v.NodeTag = "Title"
	return &v
}
func CreateLanguage() *Language {
	v := Language{}
	v.NodeTag = "Language"
	return &v
}
func CreateWidth() *Width {
	v := Width{}
	v.NodeTag = "Width"
	return &v
}
func CreateHeight() *Height {
	v := Height{}
	v.NodeTag = "Height"
	return &v
}
func CreateType() *Type {
	v := Type{}
	v.NodeTag = "Type"
	return &v
}
func CreateUpdatePeriod() *UpdatePeriod {
	v := UpdatePeriod{}
	v.NodeTag = "UpdatePeriod"
	return &v
}
func CreateUpdateFrequency() *UpdateFrequency {
	v := UpdateFrequency{}
	v.NodeTag = "UpdateFrequency"
	return &v
}
func CreateCopyright() *Copyright {
	v := Copyright{}
	v.NodeTag = "Copyright"
	return &v
}
