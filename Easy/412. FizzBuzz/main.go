/*

Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

*/

package main

import (
	"log"
	"reflect"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func main() {
	testCases := []struct {
		description string
		input       int
		expect      []string
	}{
		{
			description: "Test Case 1",
			input:       15,
			expect:      []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, t := range testCases {
		result := fizzBuzz(t.input)
		if !reflect.DeepEqual(t.expect, result) {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
