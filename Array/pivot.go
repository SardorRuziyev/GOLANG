package main

import "fmt"

func pivotIndex(nums []int) int {

	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {

		if leftSum == totalSum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("Pivot Index:", pivotIndex(nums))
}
