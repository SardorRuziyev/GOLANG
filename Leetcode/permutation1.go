package main

import "fmt"

func buildArray(num []int) []int {
	n := len(num)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = num[num[i]]

	}
	return ans
}
func main() {
	nums := []int{6, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums))
}
