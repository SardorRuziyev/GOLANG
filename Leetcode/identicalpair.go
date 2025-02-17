package main

import "fmt"

func identicalPair(nums []int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(identicalPair(nums))

}
