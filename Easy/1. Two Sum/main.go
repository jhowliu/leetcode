/*
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,

	return [0, 1].

*/
package main

import (
	"log"
	"reflect"
)

/*
func twoSum(nums []int, target int) []int {
	// Need Sorting first Time: O(nlogn)
        left, right := 0, len(nums)-1
	nums = sort.Ints(nums)
        for left < right {
                if target == nums[left] + nums[right] {
                        break
                }

                if nums[right] > target {
                        right--
                } else if nums[left] < target - nums[right] {
                        left++
                } else {
                        right--
                }
        }

        result := []int{nums[left], nums[right]}

        return result
}
*/

// using hash table time: O(N)
func twoSum(nums []int, target int) []int {
	result := []int{}

	t := map[int]int{}

	for i, num := range nums {
		// Current number find itself in the hash table if it has been store in the table with other number.
		if val, ok := t[num]; ok {
			result = append(result, val, i)
		}
		// store current index for diff number
		t[target-num] = i
	}

	return result
}

func main() {
	testCases := []struct {
		description string
		input       []int
		target      int
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       []int{2, 7, 11, 15},
			target:      9,
			expect:      []int{0, 1},
		},
		{
			description: "Test Case 1",
			input:       []int{3, 2, 4},
			target:      6,
			expect:      []int{1, 2},
		},
	}

	for _, t := range testCases {
		if result := twoSum(t.input, t.target); !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: expect:%v != actual:%v",
				t.description, t.expect, result,
			)
		}
	}
}
