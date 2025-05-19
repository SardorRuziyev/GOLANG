package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := "data.txt"

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]byte, 909)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Println("/n/n/n/n")
	data1 := make([]byte, 100)
	_, err = file.Read(data1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data1))
}
