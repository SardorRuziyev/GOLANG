package main

import "fmt"

type Car struct {
	Brand   string
	Model   string
	Year    int
	MileAge int
}

func (c *Car) Drive(km int) {
	c.MileAge += km

}

func Drive1(c *Car, km int) {
	c.MileAge += km
}

func main() {
	c := Car{
		Brand:   "Brand",
		Model:   "Model",
		Year:    2021,
		MileAge: 20,
	}

	c.Drive(34)
	fmt.Println(c)

	Drive1(&c, 12)
	fmt.Println(c)
}
