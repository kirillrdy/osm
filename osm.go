package osm

import (
	"compress/bzip2"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

type Osm struct {
	XMLName  xml.Name    `xml:"osm"`
	Node     []*Node     `xml:"node"`
	Way      []*Way      `xml:"way"`
	Relation []*Relation `xml:"relation"`
}

// Loads packaged version of Melbourne OSM ( perhaps somewhat outdated )
func LoadPackagedMelbourne() *Osm {
	return LoadFromBz2(packageDir() + "/melbourne.osm.bz2")
}

// Loads bzip'ed xml
func LoadFromBz2(filename string) *Osm {
	bzip_file, err := os.Open(filename)
	defer bzip_file.Close()
	if err != nil {
		log.Panic(err)
	}

	osm_file := bzip2.NewReader(bzip_file)

	decoder := xml.NewDecoder(osm_file)
	osm := Osm{}
	err = decoder.Decode(&osm)
	if err != nil {
		panic(err)
	}
	return &osm

}

func LoadFromJson(filename string) *Osm {
	json_file, err := os.Open(filename)
	defer json_file.Close()
	if err != nil {
		log.Panic(err)
	}

	decoder := json.NewDecoder(json_file)
	osm := Osm{}
	err = decoder.Decode(&osm)
	if err != nil {
		panic(err)
	}
	return &osm

}

func (osm *Osm) SaveToJson(filename string) {
	json_file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer json_file.Close()

	json_encoder := json.NewEncoder(json_file)
	json_encoder.Encode(osm)
}
