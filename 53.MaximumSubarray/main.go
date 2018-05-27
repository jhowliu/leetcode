/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

*/
package main

import (
	"log"
	"math"
)

// 動態規劃問題 Maximum Subarray problem
// A: Kadane's Alogorithm: O(n)內
func maxSubArray(nums []int) int {

	max := nums[0]
	sum := nums[0]

	for i := 1; i != len(nums); i++ {
		// cur+nums[i] (tricky) 加總後結果跟目前的值進行比較,
		// 若加總結果較大則表示不需要重新計算,若目前數字較大則代表加總需要重新由此數字計算(因為最大數值可能會從這數字累加後結果)
		sum = int(math.Max(float64(sum+nums[i]), float64(nums[i])))
		// 記錄著一路上碰到的最大加總(各種可能的數字集合結果)
		max = int(math.Max(float64(sum), float64(max)))
	}

	return max
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect:      6,
		},
		{
			description: "Test Case 2",
			input:       []int{-3, -2, 1, -3},
			expect:      1,
		},
	}

	for _, t := range testCases {
		result := maxSubArray(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
