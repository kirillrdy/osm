package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)

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

	data = `
<osm version="0.6" generator="Osmosis 0.43.1">
	<node id="579269" version="14" timestamp="2011-12-02T02:56:21Z" uid="11111" user="melb_guy" changeset="10013948" lat="-37.9143529" lon="145.1268578"/>
 	<node id="579270" version="5" timestamp="2010-01-11T03:34:38Z" uid="201443" user="stevage" changeset="3593374" lat="-37.914577" lon="145.135076"/>
	<node id="579287" version="2" timestamp="2009-12-10T11:22:32Z" uid="6513" user="Glen" changeset="3340743" lat="-37.9078288" lon="145.1352312">
    	<tag k="traffic_calming" v="hump"/>
  	</node>
	<way id="4074822" version="12" timestamp="2011-11-02T11:50:44Z" uid="42429" user="42429" changeset="9721118">
		<nd ref="21578112"/>
		<nd ref="579487458"/>
		<nd ref="579487477"/>
		<nd ref="21578113"/>
		<nd ref="579487502"/>
		<nd ref="579487517"/>
		<nd ref="579487527"/>
		<nd ref="579487762"/>
		<nd ref="21578114"/>
		<nd ref="579488504"/>
		<nd ref="579495286"/>
		<nd ref="579488505"/>
		<nd ref="579507010"/>
		<nd ref="579506290"/>
		<nd ref="579494627"/>
		<nd ref="60096993"/>
		<tag k="electrified" v="no"/>
		<tag k="gauge" v="1600"/>
		<tag k="railway" v="rail"/>
		<tag k="source" v="yahoo"/>
	</way>
	<relation id="3375526" version="1" timestamp="2013-12-13T02:08:39Z" uid="79475" user="AlexOnTheBus" changeset="19424318">
		<member type="way" ref="251572321" role="outer"/>
		<member type="way" ref="48570264" role="outer"/>
		<member type="way" ref="48570263" role="outer"/>
		<member type="way" ref="251572198" role="outer"/>
		<member type="way" ref="183845711" role="outer"/>
		<member type="way" ref="184363277" role="outer"/>
		<member type="way" ref="184363276" role="outer"/>
		<tag k="admin_level" v="10"/>
		<tag k="boundary" v="administrative"/>
		<tag k="name" v="Point Wilson"/>
		<tag k="postal_code" v="3212"/>
		<tag k="type" v="boundary"/>
	</relation>
  
</osm>	
	`

	osm := Osm{}
	err = xml.Unmarshal([]byte(data), &osm)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(osm)

}
