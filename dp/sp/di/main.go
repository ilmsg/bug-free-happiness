package main

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	to           *Person
	relationship Relationship
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrentOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations, Info{parent, child, Parent})
	rs.relations = append(rs.relations, Info{child, parent, Child})
}

func main() {}
