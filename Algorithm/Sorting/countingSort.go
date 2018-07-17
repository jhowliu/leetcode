package sorting

// time complexity is O(n+k)
// n is number of element and k is range of numbers
func countingSort(nums []int, RANGE int) []int {
	counter := make([]int, RANGE+1)
	result := make([]int, len(nums))

	for _, num := range nums {
		counter[num]++
	}

	// accumulate count
	for i := 1; i <= RANGE; i++ {
		counter[i] += counter[i-1]
	}
	// sort
	for _, num := range nums {
		result[counter[num]-1] = num
		counter[num]--
	}

	return result
}
