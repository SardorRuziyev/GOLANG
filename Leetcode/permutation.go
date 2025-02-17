package main

import "fmt"

func buildArray(nums []int) []int {
	ans := []int{}

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
func main() {

	nums := []int{1, 2, 3, 0}
	fmt.Println(buildArray(nums))
}
