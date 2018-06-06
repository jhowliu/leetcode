/*

Given a singly linked list, determine if it is a palindrome.

Example 1:

	Input: 1->2
	Output: false

Example 2:

	Input: 1->2->2->1
	Output: true

Follow up:
	Could you do it in O(n) time and O(1) space? Impossible to use O(1) space
*/

package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	tail := reverseList(head)

	for head != nil {
		if tail.Val != head.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	next := head.Next
	tail := &ListNode{head.Val, nil}

	for next != nil {
		tmp := ListNode{next.Val, tail}
		tail = &tmp
		next = next.Next
	}

	return tail
}

func main() {
	testCase := []struct {
		description string
		input       *ListNode
		expect      bool
	}{
		{
			description: "Test Case 1",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			expect: false,
		},
		{
			description: "Test Case 2",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			expect: true,
		},
	}

	for _, t := range testCase {
		result := isPalindrome(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
