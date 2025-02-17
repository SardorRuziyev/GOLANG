package main

import "fmt"

func calstring(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		dif := int(s[i] - s[i-1])
		if dif < 0 {
			dif = -dif
		}
		score = score + dif
	}
	return score
}
func main() {
	s := "sar"
	fmt.Println(calstring(s))

}
