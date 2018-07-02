package Sorting

import (
	"log"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {

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
	}

	for _, tc := range testCases {
		if heapSort(tc.input); !reflect.DeepEqual(tc.input, tc.expect) {
			log.Fatalf(
				"%s: expect:[%v] != result:[%v]",
				tc.description, tc.expect, tc.input,
			)
		}
	}
}
