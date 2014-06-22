package osm

import (
	"compress/bzip2"
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"path"
	"runtime"
)

func HasTagByKeyValue(item Taged, key, value string) bool {
	for _, tag := range item.Tags() {
		if tag.Key == key && tag.Value == value {
			return true
		}
	}
	return false
}

func HasTagByValue(item Taged, name string) bool {
	for _, tag := range item.Tags() {
		if tag.Value == name {
			return true
		}
	}
	return false
}

func HasTagByKey(item Taged, name string) bool {
	for _, tag := range item.Tags() {
		if tag.Key == name {
			return true
		}
	}
	return false
}

func (way *Way) IsRailWay() bool {
	return HasTagByKeyValue(way, "railway", "rail")
}

func (relation *Relation) IsTrainRoute() bool {
	return HasTagByKeyValue(relation, "route", "train")
}

func loadFromBz2() *Osm {
	bzip_file, err := os.Open("melbourne.osm.bz2")
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

func packageDir() string {
	_, current_file, _, _ := runtime.Caller(0)
	package_dir := path.Dir(current_file)
	return package_dir
}

func LoadFromJson() *Osm {
	json_file, err := os.Open(packageDir() + "/melbourne.json")
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

func saveJson(osm *Osm) {
	json_file, err := os.Create("melbourne.json")

	if err != nil {
		panic(err)
	}

	defer json_file.Close()

	json_encoder := json.NewEncoder(json_file)
	json_encoder.Encode(osm)
}

//func main() {
//
//	//osmOriginal := loadFromBz2()
//	//saveJson(osmOriginal)
//	osm := loadFromJson()
//
//	//For fast lookup
//	nodes := map[uint64]Node{}
//	ways := map[uint64]Way{}
//	relations := map[uint64]Relation{}
//
//	var frankstone_line_id uint64 = 344911
//
//	for _, node := range osm.Node {
//		nodes[node.Id] = node
//	}
//
//	for _, way := range osm.Way {
//		ways[way.Id] = way
//	}
//
//	for _, relation := range osm.Relation {
//		relations[relation.Id] = relation
//	}
//
//	frankstone_line := relations[frankstone_line_id]
//
//	first_way_id := frankstone_line.Member[0].Ref
//	first_way := ways[first_way_id]
//	fmt.Println(first_way)
//
//	for _, nd := range first_way.Nd {
//		node := nodes[nd.Ref]
//		fmt.Printf("%f %f\n", node.Lat, node.Lon)
//
//	}
//
//}
