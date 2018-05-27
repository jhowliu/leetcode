package main

import (
	"log"
)

func climbStairs(n int) int {
	// Dynamic programming
	// recursion version takes too long

	if n <= 2 {
		return n
	}

	pre1 := 2
	pre2 := 1
	count := 0

	for i := 2; i != n; i++ {
		count = pre1 + pre2
		pre2 = pre1
		pre1 = count
	}

	return count
}

func main() {
	testCases := []struct {
		description string
		input       int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       2,
			expect:      2,
		},
		{
			description: "Test Case 2",
			input:       39,
			expect:      102334155,
		},
		{
			description: "Test Case 3",
			input:       10,
			expect:      89,
		},
		{
			description: "Test Case 4",
			input:       3,
			expect:      3,
		},
	}

	for _, t := range testCases {
		result := climbStairs(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: input[%v], expect[%v] != actual[%v]",
				t.description, t.input, t.expect, result,
			)
		}
	}
}
