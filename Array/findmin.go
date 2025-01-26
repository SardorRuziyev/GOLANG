package main

import "fmt"

func main() {
	nums := []int{4, 5, 5}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	minValue := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[0] < minValue {
			minValue = nums[i]
		}
	}
	return minValue
}
