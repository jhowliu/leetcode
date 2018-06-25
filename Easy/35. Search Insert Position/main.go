/*

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
	Input: [1,3,5,6], 5
	Output: 2

Example 2:
	Input: [1,3,5,6], 2
	Output: 1

*/

package main

import (
	"log"
)

// use binary search tree. time complexity: O(logN)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left // mid+1
}

func main() {
	testCases := []struct {
		description string
		input       []int
		target      int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 3, 5, 6},
			target:      5,
			expect:      2,
		},
		{
			description: "Test Case 2",
			input:       []int{1, 3, 5, 6},
			target:      2,
			expect:      1,
		},
	}

	for _, t := range testCases {
		result := searchInsert(t.input, t.target)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
