/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]

*/
package main

import (
	"log"
	"math"
	"reflect"
)

func findDuplicates(nums []int) []int {
	result := []int{}

	for i := 0; i != len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		} else {
			// the second time, the nums[index] will be minus
			result = append(result, index+1)
		}

	}

	return result
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       []int{4, 3, 2, 7, 8, 2, 3, 1},
			expect:      []int{2, 3},
		},
		{
			description: "Test Case 2",
			input:       []int{1, 1},
			expect:      []int{1},
		},
		{
			description: "Test Case 3",
			input:       []int{1},
			expect:      []int{},
		},
	}

	for _, t := range testCases {
		result := findDuplicates(t.input)
		if !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: expect:%v != actual:%v",
				t.description, t.expect, result,
			)
		}
	}
}
