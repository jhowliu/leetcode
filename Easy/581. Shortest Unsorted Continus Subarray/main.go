/*

Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
	Input: [2, 6, 4, 8, 10, 9, 15]
	Output: 5
	Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
	Then length of the input array is in range [1, 10,000].
	The input array may contain duplicates, so ascending order here means <=.

*/

package main

import (
	"log"
)

func Min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func findUnsortedSubarray(nums []int) int {
	begin, end, n := -1, -2, len(nums)
	min, max := nums[len(nums)-1], nums[0]

	i := 0

	for i != n {
		max = Max(max, nums[i])
		min = Min(min, nums[n-i-1])

		// 由前往後找出不超過max的終點 (因為越來越大)
		if max > nums[i] {
			end = i
		}
		// 由後往前找出超過min的起始點 (因為越來越小)
		if min < nums[n-i-1] {
			begin = n - i - 1
		}
		i++
	}

	return end - begin + 1
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{2, 6, 4, 8, 10, 9, 15},
			expect:      5,
		},
		{
			description: "Test Case 2",
			input:       []int{1, 2, 3, 4},
			expect:      0,
		},
	}

	for _, t := range testCases {
		result := findUnsortedSubarray(t.input)

		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
