package SquareSum

import "testing"

func TestSquareSum(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expect      bool
	}{
		{
			description: "Test Case 1",
			input:       5,
			expect:      true,
		},
		{
			description: "Test Case 2",
			input:       99999999,
			expect:      false,
		},
		{
			description: "Test Case 3",
			input:       2,
			expect:      true,
		},
		{
			description: "Test Case 4",
			input:       0,
			expect:      true,
		},
	}

	for _, tc := range testCases {
		if result := judgeSquareSum(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: except[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
