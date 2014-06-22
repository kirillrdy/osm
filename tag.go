package osm

import "encoding/xml"

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

type Taged interface {
	Tags() []Tag
}
