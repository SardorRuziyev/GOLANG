package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	j := removeElement(nums, val)
	fmt.Println(nums[:j])

}
func removeElement(nums []int, val int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j

}
