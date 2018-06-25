/*

Given a string, find the length of the longest substring without repeating characters.

Examples:

	- Given "abcabcbb", the answer is "abc", which the length is 3.

	- Given "dvdf", the answer is "vdf", with the length of 3.

	- Given "pwwkew", the answer is "wke", with the length of 3.
	  (Note that the answer must be a substring, "pwke" is a subsequence and not a substring.)

*/
package main

import "log"

func findLongestSubstring(s string) int {
	countTable := map[byte]int{}

	left, right, max := 0, 0, 0

	for right < len(s) {
		countTable[s[right]]++

		// 遇到重複字元
		if countTable[s[right]] > 1 {
			if right-left > max {
				max = right - left
			}
			// 把指針移到重複字元第一次出現位置後一個，確保目前left-right沒有重複字元
			for left < right {
				countTable[s[left]]--

				if countTable[s[left]] == 1 {
					left++ // 第一個重複字元後一位
					break
				}
				left++
			}
		}

		right++
	}

	if right-left > max {
		max = right - left
	}

	return max
}

func main() {
	testCases := []struct {
		description string
		input       string
		expect      int
	}{
		{
			description: "Test Case 1",
			input:       "abcabcbb",
			expect:      3,
		},
		{
			description: "Test Case 2",
			input:       "dvdf",
			expect:      3,
		},
		{
			description: "Test Case 3",
			input:       "pwwkew",
			expect:      3,
		},
	}

	for _, t := range testCases {
		result := findLongestSubstring(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
