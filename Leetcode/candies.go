package main

// func main() {
// 	candies := []int{2, 3, 5, 1, 3}
// 	extraCandies := 3

// 	fmt.Println(kidsWithCandies(candies, extraCandies))

// }
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0

	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	var result []bool

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCandies {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
