package main

import "fmt"

func main() {
	nums := []int{5, 6, 14, 0}
	fmt.Println(dominantIndex(nums), nums) // Output: 1
}

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	largestIndex := 0
	// Find the index of the largest number
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[largestIndex] {
			largestIndex = i
		}
	}

	// Verify if the largest number is at least twice every other number

	sum := 0
	for i := 0; i < len(nums); i++ {

		if i != largestIndex && nums[largestIndex] >= nums[i]*2 {
			sum += 1
		}

	}

	if sum == len(nums)-1 {
		return largestIndex
	}

	return -1

}
