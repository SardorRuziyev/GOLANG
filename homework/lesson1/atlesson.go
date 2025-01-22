package main

import "fmt"

func main() {
	a, err := div(34, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cant")
	}
	c := a / b
	return c, nil
}
