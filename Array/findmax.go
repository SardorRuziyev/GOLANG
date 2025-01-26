package main

import "fmt"

func main() {
	nums := []int{3, 4, 6, 7, 9}
	fmt.Println(findMax(nums))

}

func findMax(nums []int) int {

	maxValue := nums[0]

	for i := 0; i < len(nums); i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]

		}
	}
	return maxValue
}
