package osm

import "encoding/xml"

type Member struct {
	XMLName xml.Name `xml:"member"`
	Type    string   `xml:"type,attr"`
	Ref     uint64   `xml:"ref,attr"`
	Role    string   `xml:"role,attr"`
}

type Node struct {
	XMLName xml.Name `xml:"node"`
	Id      uint64   `xml:"id,attr"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Tag     []Tag    `xml:"tag"`
}
type Nd struct {
	XMLName xml.Name `xml:"nd"`
	Ref     uint64   `xml:"ref,attr"`
}
