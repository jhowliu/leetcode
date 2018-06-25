/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.
*/
package main

import (
	"log"
)

// 因為最多數的數字數量大於長度一半，所以Majority不管最後如何，都只會剩下最多數的數字
func majorityElement(nums []int) int {
	count := 1
	majority := nums[0]

	for _, num := range nums[1:] {
		if majority == num {
			count += 1
		} else if count == 0 {
			count += 1
			majority = num
		} else {
			count -= 1
		}
	}

	return majority
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{1, 1, 2, 1, 1},
			expect:      1,
		},
		{
			description: "Test Case 2",
			input:       []int{2, 2, 2, 3, 3, 2},
			expect:      2,
		},
	}

	for _, t := range testCases {
		result := majorityElement(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != result:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
