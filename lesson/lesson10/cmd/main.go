package main

import (
	"fmt"
	"github.com/google/uuid"
	"lesson10/test"
)

func main() {
	fmt.Println(uuid.NewString())

	r := test.MakeTest("a", "b")
	fmt.Println(r)
}
