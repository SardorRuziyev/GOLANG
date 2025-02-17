package main

import "fmt"

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2}
	k := 3
	multiplier := 2
	fmt.Println(getFinalState(nums, k, multiplier))

}
func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		mini := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[mini] {
				mini = j
			}

		}

		nums[mini] *= multiplier
	}
	return nums

}
