package main

import "fmt"

type Car struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
}

func (c Car) Drive() int {
	c.Mileage 
	fmt.Print()

}

func main() {
	car := Car{
		Brand:   "Kia",
		Model:   "Rio",
		Year:    2013,
		Mileage: 160000,
	}

}
