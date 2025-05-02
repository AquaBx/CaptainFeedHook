package xml

import (
	"CaptainFeedHook/utils"
	"encoding/xml"
	"fmt"
	"time"
)

type Date struct {
	Node
	Value int64
}

type Updated = Date
type Published = Date

func CreateDate() *Date {
	v := Date{}
	v.NodeTag = "Date"
	return &v
}

func CreateUpdated() *Updated {
	v := Updated{}
	v.NodeTag = "Updated"
	return &v
}

func CreatePublished() *Published {
	v := Published{}
	v.NodeTag = "Published"
	return &v
}

func (t *Date) SetChardata(d xml.CharData) {
	for _, format := range []string{time.RFC822, time.RFC3339, time.RFC1123, time.RFC1123Z, "Mon, 02 Jan 06 15:04:05 -0700"} {
		tt, err := time.Parse(format, string(d))
		if err != nil {
		} else {
			t.Value = tt.UTC().Unix()
			return
		}
	}
	utils.Log("error", fmt.Sprintf("Cannot parse date %s", d))
}
