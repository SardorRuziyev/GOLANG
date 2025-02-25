package main

import (
	"fmt"
)

func main() {

	//fmt.Printf("%v - %v", "Hello", 22435)
	//s := fmt.Sprintf("%v - %v", "Helo", 22435)

	//r, err := div(23, 0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(r)

	//fmt.Println(arr, len(arr), cap(arr))

	//slc := make([]string, 10)
	//slc[9] = "sadasd"
	//slc = append(slc, "hello")
	//slc = append(slc, "world")
	//slc = append(slc, "world")
	//slc = append(slc, "world")
	//fmt.Println(slc, len(slc), cap(slc))

	m := map[string]string{}
	m["hello"] = "world"
	m["helo"] = "something"
}

func div(a, b int) (int, error) {
	if b == 0 {
		err := fmt.Errorf("0ga bo'lib bo'lmaydi")
		return 0, err
	}

	return a / b, nil
}
