package main

import "fmt"

func runningSum(nums []int) []int {
	answer := []int{}

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		answer = append(answer, sum)
	}

	return answer
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
