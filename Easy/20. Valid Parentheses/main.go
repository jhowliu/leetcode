/*

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
	- Open brackets must be closed by the same type of brackets.
	- Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:
	Input: "()"
	Output: true
*/

package main

import (
	"log"
)

func isValid(s string) bool {
	stack := []int32{}

	for _, c := range s {
		if c == int32('{') || c == int32('[') || c == int32('(') {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			element := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if element == int32('{') && c != int32('}') {
				return false
			}
			if element == int32('[') && c != int32(']') {
				return false
			}
			if element == int32('(') && c != int32(')') {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func main() {
	testCases := []struct {
		description string
		input       string
		expect      bool
	}{
		{
			description: "Test Case 1",
			input:       "{}",
			expect:      true,
		},
		{
			description: "Test Case 2",
			input:       "[",
			expect:      false,
		},
		{
			description: "Test Case 3",
			input:       "",
			expect:      true,
		},
	}

	for _, t := range testCases {
		result := isValid(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
