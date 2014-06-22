package osm

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
