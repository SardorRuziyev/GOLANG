package main

import (
	"fmt"
)

func main() {
	//res := add(1, 2)
	//fmt.Println(res)
	//
	////a, b, q, err := div(423, 1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(a, b, q)
	//
	//fmt.Println(just())
	a, b, err := calc()
	if err != nil {
		panic(err)
	}
	fmt.Println(a, b)
}

func add(a, b int) int {
	c := a + b
	return c
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("we cant div to zero(0)")
	}
	c := a / b
	return c, nil
}

func calc() (int, int, error) {
	var q, w int
	fmt.Scan(&q, &w)
	added := add(q, w)

	dived, err := div(q, w)
	if err != nil {
		return 0, 0, err
	}

	return added, dived, nil
}

func just() (bool, bool, bool, bool, bool, bool, string, int) {
	return false, true, true, true, false, true, "fadsfdasjfk", 2342342453
}
