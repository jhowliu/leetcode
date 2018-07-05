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
			input:       "0001111",
			expect:      3,
		},
		{
			description: "Test Case 2",
			input:       "000111000111",
			expect:      9,
		},
		{
			description: "Test Case 3",
			input:       "",
			expect:      1,
		},
		{
			description: "Test Case 4",
			input:       "1",
			expect:      0,
		},
	}

	for _, tc := range testCases {
		if result := countBinarySubstrings(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
