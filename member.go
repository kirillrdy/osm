package osm

import "encoding/xml"

type Member struct {
	XMLName xml.Name `xml:"member"`
	Type    string   `xml:"type,attr"`
	Ref     uint64   `xml:"ref,attr"`
	Role    string   `xml:"role,attr"`
}
