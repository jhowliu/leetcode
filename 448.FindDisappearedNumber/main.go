/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/
package main

import (
	"log"
	"math"
	"reflect"
)

func findDisappearedNumbers(nums []int) []int {
	result := []int{}

	for i := 0; i != len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		// the index has not been marked if nums[index] > 0
		if nums[index] > 0 {
			// we could never change the original value using minus
			nums[index] = -nums[index]
		}
	}

	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
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
			expect:      []int{5, 6},
		},
		{
			description: "Test Case 2",
			input:       []int{1, 1},
			expect:      []int{2},
		},
		{
			description: "Test Case 3",
			input:       []int{1},
			expect:      []int{},
		},
	}

	for _, t := range testCases {
		result := findDisappearedNumbers(t.input)
		if !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: expect:%v != actual:%v",
				t.description, t.expect, result,
			)
		}
	}
}
