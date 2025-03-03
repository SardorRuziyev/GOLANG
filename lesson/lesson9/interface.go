package main

import (
	"fmt"
)

type interface1 interface {
	Name(string) Person
}
type s int

func main() {
	p := Person{}
	n := newS{}

	//var i interface1 = p
	//var i1 interface1 = n

	ExecName(p)
	ExecName(n)
}

type Person struct {
	name   string
	second string
}

func (p Person) Name(name string) Person {
	p.name = name
	fmt.Printf("in NAME 1 %p %s\n", &p, p.name)
	return p
}

type newS struct{}

func (n newS) Name(name string) Person {
	return Person{}
}

func (n newS) Second(second string) Person {
	return Person{}
}

func ExecName(i interface1) {
	i.Name("asdf")
}
