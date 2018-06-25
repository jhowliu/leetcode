/*
(Dynamic Programming)

On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.
*/
package main

import (
	"log"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)

	dp[0] = cost[0]
	dp[1] = cost[1]

	// Bottom-Up Dynamic Programming
	for i := 2; i != n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 2, 2, 2},
			expect:      3,
		},
		{
			description: "Test Case 1",
			input:       []int{1, 2, 2, 0},
			expect:      2,
		},
	}

	for _, t := range testCases {
		result := minCostClimbingStairs(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
