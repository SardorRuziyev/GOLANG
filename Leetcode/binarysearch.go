package main

import "fmt"

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for high >= low {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(binarySearch(nums, target))

}
