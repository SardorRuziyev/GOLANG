package main

import "fmt"

func duplicate(nums []int) bool {
	mp := map[int]bool{}
	mp := map[int]bool{}
	for _, v := range nums {
		if mp[nums[v]] {
			return true
		}
		mp[nums[v]] = true
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 4, 1}
	fmt.Println(duplicate(nums))

}
