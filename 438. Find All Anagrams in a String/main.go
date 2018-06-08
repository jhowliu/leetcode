/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:
	Input:
		s: "cbaebabacd" p: "abc"
	Output:
		[0, 6]
	Explanation:
		The substring with start index = 0 is "cba", which is an anagram of "abc".
		The substring with start index = 6 is "bac", which is an anagram of "abc".

*/

package main

import (
	"log"
	"reflect"
)

func findAnagrams(s string, p string) []int {
	results := []int{}

	countTable := map[byte]int{}

	for _, c := range p {
		countTable[byte(c)]++
	}

	left, right, count := 0, 0, len(p)

	for right < len(s) {

		countTable[s[right]]--
		// 若小於0表示這個字元要麻不存在於p，要麻存在但重複出現
		if countTable[s[right]] >= 0 {
			count--
		}

		if count == 0 {
			results = append(results, left)
		}

		if (right - left + 1) == len(p) {
			countTable[s[left]]++
			// 避免未出現的字元或重複字元擾亂count
			if countTable[s[left]] > 0 {
				count++
			}
			left++
		}

		right++
	}

	return results
}

func main() {
	testCases := []struct {
		description string
		input       string
		target      string
		expect      []int
	}{
		{
			description: "Test Case 1",
			input:       "cbaebabacd",
			target:      "abc",
			expect:      []int{0, 6},
		},
		{
			description: "Test Case 2",
			input:       "abab",
			target:      "ab",
			expect:      []int{0, 1, 2},
		},
	}

	for _, t := range testCases {
		result := findAnagrams(t.input, t.target)
		if !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}

}
