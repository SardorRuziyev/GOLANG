package main

import "fmt"

type Car struct {
	Brand   string
	Model   string
	Year    int
	Mileage int
}

func (c Car) Drive(km int) int {
	c.Mileage += km
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
