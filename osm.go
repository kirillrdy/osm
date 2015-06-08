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

//TODO move elsewhere
// Used for fast querying
type OsmIndex struct {
	// these are for indexing
	nodesById    map[uint64]*Node
	waysById     map[uint64]*Way
	relationById map[uint64]*Relation
}

func (osm *Osm) BuildIndex() OsmIndex {

	index := OsmIndex{}

	index.nodesById = map[uint64]*Node{}
	index.waysById = map[uint64]*Way{}
	index.relationById = map[uint64]*Relation{}

	for i := range osm.Node {
		index.nodesById[osm.Node[i].Id] = osm.Node[i]
	}
	for i := range osm.Way {
		index.waysById[osm.Way[i].Id] = osm.Way[i]
	}
	for i := range osm.Relation {
		index.relationById[osm.Relation[i].Id] = osm.Relation[i]
	}
	return index
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

func (index *OsmIndex) NodeById(id uint64) *Node {
	return index.nodesById[id]
}

func (index *OsmIndex) WayById(id uint64) *Way {
	return index.waysById[id]
}

func (index *OsmIndex) RelationById(id uint64) *Relation {
	return index.relationById[id]
}
