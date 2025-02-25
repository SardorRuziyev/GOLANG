package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error

	r, err := div(1, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
