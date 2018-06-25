/*
	Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

	The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

	Note:

	Your returned answers (both index1 and index2) are not zero-based.
	You may assume that each input would have exactly one solution and you may not use the same element twice.
	Example:

	Input: numbers = [2,7,11,15], target = 9
	Output: [1,2]

	Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/
package main

import (
	"log"
	"reflect"
)

// Space Complexity: O(1), Time Complexity: O(n)
func twoSum(numbers []int, target int) []int {
	result := []int{}

	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			result = append(result, l+1, r+1)
			break
		}

		if numbers[r] > target {
			r--
		} else if target-numbers[r] > numbers[l] {
			l++
		} else {
			r--
		}
	}

	return result
}

// Space Complexity: O(n-1), Time Complexity: O(n)
func twoSumWithHash(numbers []int, target int) []int {
	result := []int{}

	hashTable := map[int]int{}

	for i, num := range numbers {
		if _, ok := hashTable[num]; ok {
			result = append(result, hashTable[num], i+1)
			break
		}
		hashTable[target-num] = i + 1
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
			expect:      []int{1, 2},
		},
		{
			description: "Test Case 1",
			input:       []int{2, 3, 4},
			target:      6,
			expect:      []int{1, 3},
		},
	}

	for _, t := range testCases {
		if result := twoSum(t.input, t.target); !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: expect:%v != actual:%v",
				t.description, t.expect, result,
			)
		}

		if result := twoSumWithHash(t.input, t.target); !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"(TwoSumWithHash) %s: expect:%v != actual:%v",
				t.description, t.expect, result,
			)
		}
	}
}
