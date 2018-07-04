/*
We define a harmonious array is an array where the difference between its maximum value and its minimum value is exactly 1.

Example 1:
	Input: [1,3,2,2,5,2,3,7]
	Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].
Note: The length of the input array will not exceed 20,000.

*/
package solution

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 2, 3, 4, 5},
			expect:      2,
		},
		{
			description: "Test Case 2",
			input:       []int{1, 1, 1, 1},
			expect:      0,
		},
		{
			description: "Test Case 3",
			input:       []int{1, 3, 2, 2, 5, 2, 3, 7},
			expect:      5,
		},
		{
			description: "Test Case 4",
			input:       []int{1},
			expect:      0,
		},
	}

	for _, tc := range testCases {
		if result := findLHS(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}

}
