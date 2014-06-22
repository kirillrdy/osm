package osm

import "encoding/xml"

type Relation struct {
	XMLName xml.Name `xml:"relation"`
	Id      uint64   `xml:"id,attr"`
	Tag     []Tag    `xml:"tag"`
	Member  []Member `xml:"member"`
}

func (relation *Relation) IsTrainRoute() bool {
	return HasTagByKeyValue(relation, "route", "train")
}

func (relation Relation) NodeMembers() []Member {
	var members []Member
	for _, member := range relation.Member {
		if member.Type == "node" {
			members = append(members, member)
		}
	}
	return members
}

func (relation Relation) WayMembers() []Member {
	var members []Member
	for _, member := range relation.Member {
		if member.Type == "way" {
			members = append(members, member)
		}
	}
	return members
}

func (relation Relation) Tags() []Tag {
	return relation.Tag
}
