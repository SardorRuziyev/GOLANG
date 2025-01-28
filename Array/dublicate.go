package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	j := removeDuplicate(nums)
	fmt.Println("Modified array:", nums[:j])
	fmt.Println("Number of unique elements:", nums)

}

func removeDuplicate(nums []int) int {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}
	j++
}
