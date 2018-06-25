package main

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	testCases := []struct {
		description string
		input       Points
		expect      int
	}{
		{
			description: "Test case 1",
			input:       [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			expect:      2,
		},
		{
			description: "Test case 2",
			input:       [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {1, 2}},
			expect:      2,
		},
		{
			description: "Test case 3",
			input:       [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {1, 2}, {5, 6}},
			expect:      3,
		},
	}

	for _, tc := range testCases {
		result := FindMinArrowShots(tc.input)
		if result != tc.expect {
			t.Fatalf(
				"%s failed: Expect[%v] != Result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
