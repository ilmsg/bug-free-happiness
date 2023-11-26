package main

type Relationship int

const (
	Ask Relationship = iota
	Accept
	Breakup
)

type Relationships struct {
	relations []*Info
}

// a ask b for friend
// b accept a for fiend
// b breakup a for friend

func (rs *Relationships) Add() {}

type Info struct {
	from         *Person
	to           *Person
	relationship Relationship
}

type Person struct {
	name string
}

func main() {

}
