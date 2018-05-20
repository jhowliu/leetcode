package main

import (
	"log"
	"reflect"
)

/* Time Complexity: O(n^2) */
func moveZeros(nums []int) []int {
	i := 0
	count := 0

	for {
		if i >= len(nums)-count {
			break
		}

		if nums[i] == 0 {
			j := i
			for j != len(nums)-1 {
				nums[j] = nums[j+1]
				j += 1
			}
			nums[j] = 0
			count += 1
		} else {
			i += 1
		}
	}
	return nums
}

/* Time Complexity: O(n) */
func BestAnswer(nums []int) []int {
	insertZeroPos := 0

	// move all non-zero numbers forword, find the index to push zeroes.
	for _, num := range nums {
		if num != 0 {
			nums[insertZeroPos] = num
			insertZeroPos += 1
		}
	}

	// fill out zeroes
	for insertZeroPos < len(nums) {
		nums[insertZeroPos] = 0
		insertZeroPos += 1
	}

	return nums
}

func main() {
	testCases := []struct {
		description string
		input       []int
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       []int{0, 1, 0, 3, 12},
			expect:      []int{1, 3, 12, 0, 0},
		},
		{
			description: "Test Case 2",
			input:       []int{1},
			expect:      []int{1},
		},
		{
			description: "Test Case 3",
			input:       []int{0},
			expect:      []int{0},
		},
	}

	for _, t := range testCases {
		result := moveZeros(t.input)
		if !reflect.DeepEqual(t.expect, result) {
			log.Fatalf(
				"%s: expect: %v != actual: %v",
				t.description, t.expect, result,
			)
		}
	}
}
