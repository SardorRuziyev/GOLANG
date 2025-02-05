package main

import "fmt"

func main() {

	fmt.Println(getConcatenation([]int{1, 2, 4}))

}
func getConcatenation(nums []int) []int {
	ans := nums
	ans = append(ans, nums...)

	return ans

}
