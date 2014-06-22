package main

import (
	"compress/bzip2"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func (way *Way) HasTagByKeyValue(key, value string) bool {
	for _, tag := range way.Tag {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
}

func (way *Way) HasTagByValue(name string) bool {
	for _, tag := range way.Tag {
		if tag.Value == name {
			return true
		}
	}
	return false
}

func (way *Way) HasTagByKey(name string) bool {
	for _, tag := range way.Tag {
		if tag.Key == name {
			return true
		}
	}
	return false
}

func (way *Way) IsRailWay() bool {
	return way.HasTagByKeyValue("railway", "rail")
}

func main() {
	bzip_file, err := os.Open("melbourne.osm.bz2")
	if err != nil {
		log.Panic(err)
	}

	osm_file := bzip2.NewReader(bzip_file)

	decoder := xml.NewDecoder(osm_file)

	//For fast lookup
	nodes := map[uint64]Node{}
	ways := map[uint64]Way{}
	relations := map[uint64]Relation{}

	osm := Osm{}
	err = decoder.Decode(&osm)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	var frankstone_line uint64 = 344911

	for _, node := range osm.Node {
		nodes[node.Id] = node
	}

	for _, way := range osm.Way {
		ways[way.Id] = way
	}

	for _, relation := range osm.Relation {
		relations[relation.Id] = relation
	}

	fmt.Println(relations[frankstone_line].Id)
	fmt.Println(relations[frankstone_line])

}
