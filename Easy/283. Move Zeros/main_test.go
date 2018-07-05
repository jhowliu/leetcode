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
			input:       []int{0, 1, 0, 2, 3, 0},
			expect:      []int{1, 2, 3, 0, 0, 0},
		},
		{
			description: "Test Case 2",
			input:       []int{0, 0, 0, 0},
			expect:      []int{0, 0, 0, 0},
		},
		{
			description: "Test Case 3",
			input:       []int{10, 2, 0, 1, 1, 0},
			expect:      []int{10, 2, 1, 1, 0, 0},
		},
	}

	for _, tc := range testCases {
		tmp := make([]int, len(tc.input))
		copy(tmp, tc.input)
		if moveZeros(tmp); !reflect.DeepEqual(tmp, tc.expect) {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, tmp,
			)
		}
	}
}
