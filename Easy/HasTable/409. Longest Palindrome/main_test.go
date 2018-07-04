package solution

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       "bbasdlkjlascslkdasbbbbbccccc",
			expect:      25,
		},
		{
			description: "Test Case 2",
			input:       "abccccdd",
			expect:      7,
		},
		{
			description: "Test Case 3",
			input:       "asdasdasd",
			expect:      7,
		},
		{
			description: "Test Case 4",
			input:       "a",
			expect:      1,
		},
	}

	for _, tc := range testCases {
		if result := longestPalindrome(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
