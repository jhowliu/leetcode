package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	testCases := []struct {
		description string
		input       IntervalSlice
		expect      int
	}{
		{
			description: "Test Case 1",
			input: IntervalSlice{
				Interval{
					Start: 1,
					End:   100,
				},
				Interval{
					Start: 2,
					End:   50,
				},
				Interval{
					Start: 1,
					End:   10,
				},
				Interval{
					Start: 1,
					End:   2,
				},
			},
			expect: 2,
		}, {
			description: "Test Case 2",
			input: IntervalSlice{
				Interval{
					Start: 1,
					End:   2,
				},
				Interval{
					Start: 1,
					End:   2,
				},
				Interval{
					Start: 1,
					End:   2,
				},
			},
			expect: 2,
		},
	}

	for _, tc := range testCases {
		result := eraseOverlapIntervals(tc.input)
		if result != tc.expect {
			t.Fatalf(
				"%s failed: expect:[%v] != actual:[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
