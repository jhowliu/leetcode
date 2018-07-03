package sorting

import (
	"log"
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {

	testCases := []struct {
		description string
		input       []int
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       []int{2, 10, 5, 11, 20},
			expect:      []int{2, 5, 10, 11, 20},
		},
		{
			description: "Test Case 2",
			input:       []int{10, 100, 20, 2, 1},
			expect:      []int{1, 2, 10, 20, 100},
		},
		{
			description: "Test Case 3",
			input:       []int{},
			expect:      []int{},
		},
		{
			description: "Test Case 4",
			input:       []int{1, 1, 2, 2, 3},
			expect:      []int{1, 1, 2, 2, 3},
		},
	}

	for _, tc := range testCases {
		tmp := make([]int, len(tc.input))

		copy(tmp, tc.input)
		// Heap Sort
		if heapSort(tmp); !reflect.DeepEqual(tmp, tc.expect) {
			log.Fatalf(
				"%s: expect:[%v] != result:[%v]",
				tc.description, tc.expect, tmp,
			)
		}

		copy(tmp, tc.input)
		// Qucik Sort
		if quickSort(tmp, 0, len(tc.input)-1); !reflect.DeepEqual(tmp, tc.expect) {
			log.Fatalf(
				"%s: expect:[%v] != result:[%v]",
				tc.description, tc.expect, tmp,
			)
		}
	}
}
