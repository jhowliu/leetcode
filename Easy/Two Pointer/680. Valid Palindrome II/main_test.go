package Palindrome

import "testing"

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expect      bool
	}{
		{
			description: "Test Case 1",
			input:       "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			expect:      true,
		},
		{
			description: "Test Case 2",
			input:       "abca",
			expect:      true,
		},
		{
			description: "Test Case 3",
			input:       "a",
			expect:      true,
		},
		{
			description: "Test Case 4",
			input:       "abab",
			expect:      true,
		},
	}

	for _, tc := range testCases {
		if result := validPalindrome(tc.input); result != tc.expect {
			t.Fatalf(
				"%s: expect[%v] != result[%v]",
				tc.description, tc.expect, result,
			)
		}
	}
}
