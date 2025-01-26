package main

import "fmt"

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	largest := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != largest && nums[largest] < nums[i]*2 {
			return -1
		}
	}
	return largest
}
