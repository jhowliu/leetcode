package solution

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])

	if m*n != r*c {
		return nums
	}

	// Initalize two-dimension slice
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	ix := 0

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result[i][j] = nums[ix/n][ix%n]
			ix++
		}
	}

	return result
}
