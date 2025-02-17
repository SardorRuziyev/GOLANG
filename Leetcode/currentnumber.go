package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	answer := []int{}

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if i != j && nums[j] < nums[i] {
				count++
			}
		}
		answer = append(answer, count)
	}
	return answer
}
func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))

}
