package solution

import "testing"

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input1      string
		input2      string
		expect      bool
	}{
		{
			description: "Test Case 1",
			input1:      "cat",
			input2:      "rat",
			expect:      false,
		},
		{
			description: "Test Case 2",
			input1:      "gggbbbbb",
			input2:      "bbbbbggg",
			expect:      true,
		},
		{
			description: "Test Case 3",
			input1:      "",
			input2:      "",
			expect:      true,
		},
	}

	for _, tc := range testCases {
		if result := isAnagram(tc.input1, tc.input2); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
