package solution

import "testing"

func TestFindOddNumber(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 2, 3, 5, 1, 2, 3},
			expect:      5,
		},
		{
			description: "Test Case 2",
			input:       []int{9, 3, 9, 3, 9, 9, 7},
			expect:      7,
		},
	}

	for _, tc := range testCases {
		if result := findOddOccurencesNumber(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
