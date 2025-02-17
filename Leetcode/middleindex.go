package main

import "fmt"

func main() {
	nums := []int{2, 3, -1, 8, 4}
	fmt.Println(findMiddleIndex(nums))

}
func findMiddleIndex(nums []int) int {
	totalSum := 0

	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	leftSum := 0

	for i := 0; i < len(nums); i++ {
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return nums[i]
		}
		leftSum = leftSum + nums[i]
	}
	return -1
}
