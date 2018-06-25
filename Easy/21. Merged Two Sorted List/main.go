package main

import (
	"log"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func array2LinkList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := ListNode{nums[0], nil}
	next := &head // 存指標以免後面資料bind不上去

	for _, num := range nums[1:] {
		node := ListNode{num, nil}
		next.Next = &node
		next = &node
	}

	return &head
}

func linkList2Array(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := ListNode{}
	next := &head

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			next.Next = l2
			l2 = l2.Next
		} else {
			next.Next = l1
			l1 = l1.Next
		}
		next = next.Next
	}

	// appending remaining nodes
	for l1 != nil {
		next.Next = l1
		l1 = l1.Next
		next = next.Next
	}

	// appending remaining nodes
	for l2 != nil {
		next.Next = l2
		l2 = l2.Next
		next = next.Next
	}

	return head.Next
}

func main() {
	testCases := []struct {
		description string
		l1          []int
		l2          []int
		expect      []int
	}{
		{
			description: "Test Case 1",
			l1:          []int{1, 2, 3},
			l2:          []int{1, 2, 4},
			expect:      []int{1, 1, 2, 2, 3, 4},
		},
		{
			description: "Test Case 2",
			l1:          []int{1, 2, 4},
			l2:          []int{1, 3, 4},
			expect:      []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, t := range testCases {
		result := mergeTwoLists(array2LinkList(t.l1), array2LinkList(t.l2))
		if !reflect.DeepEqual(array2LinkList(t.expect), result) {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, linkList2Array(result),
			)
		}
	}
}
