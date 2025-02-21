package main

import (
	"fmt"
	"sort"
)

func getSneakya(nums []int) []int {
	sort.Ints(nums)

	answer := make([]int, 0, 2)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			answer = append(answer, nums[i])
		}
	}
	return answer
}

func main() {

	nums := []int{0, 3, 2, 1, 3, 2}
	result := getSneakya(nums)
	fmt.Println(result)
}
