package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		description string
		input       [][]int
		dim         []int
		expect      [][]int
	}{
		{
			description: "Test Case 1",
			input:       [][]int{[]int{1, 2}, []int{3, 5}},
			dim:         []int{1, 4},
			expect:      [][]int{[]int{1, 2, 3, 4}},
		},
		{
			description: "Test Case 2",
			input:       [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			dim:         []int{3, 2},
			expect:      [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}},
		},
	}

	for _, tc := range testCases {
		if result := matrixReshape(tc.input, tc.dim[0], tc.dim[1]); reflect.DeepEqual(tc.expect, result) {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
