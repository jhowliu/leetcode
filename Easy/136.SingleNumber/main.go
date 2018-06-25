package main

import (
	"log"
)

func findSingleNumber(nums []int) int {
	result := nums[0]

	for _, num := range nums[1:] {
		result = result ^ num
	}

	return result
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 2, 3, 2, 1},
			expect:      3,
		},
		{
			description: "Test Case 2",
			input:       []int{2, 2, 1},
			expect:      1,
		},
	}

	for _, t := range testCases {
		result := findSingleNumber(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: input[%v], expect[%v] != actual[%v]",
				t.description, t.input, t.expect, result,
			)
		}
	}
}
