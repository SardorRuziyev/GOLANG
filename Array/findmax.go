package main

import "fmt"

func main() {
	nums := []int{3, 4, 6, 7, 9}
	extra := 3
	fmt.Println(findMax(nums))

}

func findMax(nums []int) int {

	maxValue := nums[0]

	for i := 0; i < len(nums); i++ {
		if maxValue < nums[i] {
			maxValue = nums[i]

		}
	}
	 for i:=0; i< len(nums); i++ {
		if extra + maxValue >= maxValue {
			result[i] = true
		} else {
			result [i]= false
		}
	 }
	return maxValue
}

func findMax(nums []int)int{

	maxValue := 0

	for i:=0; i<len(nums); i++{
		if 
	}

}
func findMax(nums []int)int {

}

func findMin(nums []int){

}

removeDip(nums []int)int{

}