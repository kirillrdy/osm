package osm

import (
	"compress/bzip2"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

type Osm struct {
	XMLName  xml.Name   `xml:"osm"`
	Node     []Node     `xml:"node"`
	Way      []Way      `xml:"way"`
	Relation []Relation `xml:"relation"`

	nodesById    map[uint64]Node
	waysById     map[uint64]Way
	relationById map[uint64]Relation
}

func (osm *Osm) BuildIndex() {

	osm.nodesById = map[uint64]Node{}
	osm.waysById = map[uint64]Way{}
	osm.relationById = map[uint64]Relation{}

	for _, node := range osm.Node {
		osm.nodesById[node.Id] = node
	}
	for _, way := range osm.Way {
		osm.waysById[way.Id] = way
	}
	for _, relation := range osm.Relation {
		osm.relationById[relation.Id] = relation
	}
}

func LoadPackagedMelbourne() *Osm {
	return LoadFromBz2(packageDir() + "/melbourne.osm.bz2")
}

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

func (osm *Osm) NodeById(id uint64) *Node {
	node := osm.nodesById[id]
	return &node

	for _, node := range osm.Node {
		if node.Id == id {
			return &node
		}
	}
	return nil
}

func (osm *Osm) WayById(id uint64) *Way {
	way := osm.waysById[id]
	return &way

	for _, way := range osm.Way {
		if way.Id == id {
			return &way
		}
	}
	return nil
}

func (osm *Osm) RelationById(id uint64) *Relation {
	relation := osm.relationById[id]

	return &relation
	for _, relation := range osm.Relation {
		if relation.Id == id {
			return &relation
		}
	}
	return nil
}
