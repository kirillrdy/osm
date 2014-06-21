package main

import "encoding/xml"

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}
type Member struct {
	XMLName xml.Name `xml:"member"`
	Type    string   `xml:"type,attr"`
	Ref     string   `xml:"ref,attr"`
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
type Way struct {
	XMLName xml.Name `xml:"way"`
	Id      uint64   `xml:"id,attr"`
	Tag     []Tag    `xml:"tag"`
	Nd      []Nd     `xml:"nd"`
}

type Relation struct {
	XMLName xml.Name `xml:"relation"`
	Id      uint64   `xml:"id,attr"`
	Tag     []Tag    `xml:"tag"`
	Member  []Member `xml:"member"`
}

type Osm struct {
	XMLName  xml.Name   `xml:"osm"`
	Node     []Node     `xml:"node"`
	Way      []Way      `xml:"way"`
	Relation []Relation `xml:"relation"`
}
