package osm

//TODO move elsewhere
// Used for fast querying
type Index struct {
	// these are for indexing
	nodesById    map[uint64]*Node
	waysById     map[uint64]*Way
	relationById map[uint64]*Relation
}

func (osm *Osm) BuildIndex() Index {

	index := Index{}

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

func (index *Index) NodeById(id uint64) *Node {
	return index.nodesById[id]
}

func (index *Index) WayById(id uint64) *Way {
	return index.waysById[id]
}

func (index *Index) RelationById(id uint64) *Relation {
	return index.relationById[id]
}
