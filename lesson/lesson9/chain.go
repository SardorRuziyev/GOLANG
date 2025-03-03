package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	t = t.Add(4234)

	p := person{}
	fmt.Printf("%p %s\n", &p, p.name)
	e := p.Name("asdf")
	fmt.Printf("%p\n %s\n", &e, e.name)

	p = Name("sdfg", p)
	fmt.Printf("%p %s\n", &p, p.name)

	p.NameP("fsdgfdf")
	fmt.Printf("%p %s\n", &p, p.name)
}

type person struct {
	name   string
	second string
}

func (p *person) NameP(name string) {
	p.name = name
	fmt.Printf("in NAMEP 3 %p %s\n", p, p.name)
}

func (p person) Name(name string) person {
	p.name = name
	fmt.Printf("in NAME 1 %p %s\n", &p, p.name)
	return p
}

func Name(name string, p person) person {
	p.name = name
	fmt.Printf(" IN NAME 2 %p %s\n", &p, p.name)
	return p
}
