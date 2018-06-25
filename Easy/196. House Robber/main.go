package main

import (
	"log"
	"math"
)

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}

	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]

	for i := 3; i != n; i++ {
		dp[i] = nums[i] + int(math.Max(float64(dp[i-2]), float64(dp[i-3])))
	}

	return int(math.Max(float64(dp[n-1]), float64(dp[n-2])))
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 2, 3, 1},
			expect:      4,
		},
		{
			description: "Test Case 1",
			input:       []int{2, 7, 9, 3, 1},
			expect:      12,
		},
	}

	for _, t := range testCases {
		result := rob(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
