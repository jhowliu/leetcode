/*
Write a program to find the node at which the intersection of two singly linked lists begins.


For example, the following two linked lists:

	A:       a1 → a2
			↘
			c1 → c2 → c3
			↗
	B:  b1 → b2 → b3

begin to intersect at node c1.

*/

package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 設 A 的長度為 a + c，B 的長度為 b + c，其中 c 為尾部公共部分長度，可知 a + c + b = b + c + a。
// 當訪問A List指針訪問到鍊錶尾部時，令它開始訪問B List；同樣地，當訪問B List的指針訪問到鍊錶尾部時，令它從開始訪問A List。
// 這樣就能控制訪問 A 和 B 兩個鍊錶的指針能同時訪問到交點。
func findIntersection(headA *ListNode, headB *ListNode) *ListNode {
	a, b := headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}

func main() {
	intersectionNode := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
		},
	}

	testCase := []struct {
		description string
		inputA      *ListNode
		inputB      *ListNode
		expect      *ListNode
	}{
		{
			description: "Test Case 1",
			inputA: &ListNode{
				Val:  2,
				Next: intersectionNode,
			},
			inputB: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  10,
					Next: intersectionNode,
				},
			},
			expect: intersectionNode,
		},
	}

	for _, t := range testCase {
		result := findIntersection(t.inputA, t.inputB)
		log.Println(result, &t.expect)
		if result != t.expect {
			log.Fatalf(
				"%s: expect[%v] != actual[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
