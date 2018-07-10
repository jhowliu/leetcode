package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 1, 2, 2, 2, 2, 3, 3, 3, 3},
			expect:      []int{3, 2},
		},
		{
			description: "Test Case 2",
			input:       []int{1, 2},
			expect:      []int{1, 2},
		},
		{
			description: "Test Case 3",
			input:       []int{},
			expect:      []int{},
		},
		{
			description: "Test Case 4",
			input:       []int{1, 1, 1, 1, 3, 3, 3},
			expect:      []int{1, 3},
		},
	}

	for _, tc := range testCases {
		if result := findMajorityElements(tc.input); !reflect.DeepEqual(tc.expect, result) {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
