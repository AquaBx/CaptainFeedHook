package xml

import (
	"testing"
)

const xmltest = `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom"
    xmlns:dc="http://purl.org/dc/elements/1.1/">
    <channel>
        <title>ChannelTitle</title>
        <link>ChannelLink</link>
        <description>ChannelDescription</description>
        <atom:link href="ChannelAtomLink" rel="self" type="application/rss+xml" />
        <item>
            <title>ItemTitle</title>
            <link>ItemLink</link>
            <summary>ItemSummary</summary>
            <description>ItemDescription</description>
            <category>ItemCategory</category>
            <guid>ItemGuid</guid>
            <dc:creator>ItemDcCreator</dc:creator>
            <pubDate>Fri, 25 Apr 2025 12:00:00 +0000</pubDate>
            <image>ItemImage</image>
        </item>
    </channel>
</rss>
`

// TestParseRSS calls xml
func TestParseRSS(t *testing.T) {
	xmlv := Visitor([]byte(xmltest))

	for _, Channel := range xmlv.Channels {
		for _, Entry := range Channel.Entries {
			if Entry.Content != "ItemDescription" {
				t.Errorf("Should log")
			}
			if Entry.Dates[0].Value != 1745582400 {
				t.Errorf("Should log")
			}
			if Entry.Title != "ItemTitle" {
				t.Errorf("Should log")

			}
			if Entry.Links[0].Href != "ItemLink" {
				t.Errorf("Should log")

			}
			if Entry.People[0].Name != "ItemDcCreator" {
				t.Errorf("Should log")

			}
			if Entry.Categories[0] != "ItemCategory" {
				t.Errorf("Should log")

			}
		}
	}
}
