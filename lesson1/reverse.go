package main

import "fmt"

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func main() {

	// Reversing the string.
	str := "Pizza"

	// returns the reversed string.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
