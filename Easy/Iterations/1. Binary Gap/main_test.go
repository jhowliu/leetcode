package solution

import (
	"log"
	"testing"
)

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       32,
			expect:      0,
		},
		{
			description: "Test Case 2",
			input:       1610612737,
			expect:      28,
		},
		{
			description: "Test Case 3",
			input:       1162,
			expect:      3,
		},
		{
			description: "Test Case 4",
			input:       561892,
			expect:      3,
		},
	}

	for _, tc := range testCases {
		if result := findBinaryGap(tc.input); result != tc.expect {
			log.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
