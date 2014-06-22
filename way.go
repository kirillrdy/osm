package osm

import "encoding/xml"

type Way struct {
	XMLName xml.Name `xml:"way"`
	Id      uint64   `xml:"id,attr"`
	Tag     []Tag    `xml:"tag"`
	Nd      []Nd     `xml:"nd"`
}

func (way *Way) IsRailWay() bool {
	return HasTagByKeyValue(way, "railway", "rail")
}

func (way Way) Tags() []Tag {
	return way.Tag
}
