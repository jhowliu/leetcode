package main

import (
	"log"
	"math"
)

func reverseInteger(input int) int {
	result := 0

	for input != 0 {
		remainer := input % 10
		result = result*10 + remainer
		input /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}

func main() {
	testCases := []struct {
		description string
		input       int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       123,
			expect:      321,
		},
		{
			description: "Test Case 2",
			input:       -123,
			expect:      -321,
		},
		{
			description: "Test Case 3",
			input:       1534236469,
			expect:      0,
		},
	}

	for _, t := range testCases {
		result := reverseInteger(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: input[%v], expect[%v] != actual[%v]",
				t.description, t.input, t.expect, result,
			)
		}
	}
}
