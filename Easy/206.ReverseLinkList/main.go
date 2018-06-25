package main

import (
	"log"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
	last := head

	for next != nil {
		tmp := next.Next
		next.Next = last
		last = next
		next = tmp
	}

	head.Next = nil

	return last
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

func main() {
	testCases := []struct {
		description string
		input       *ListNode
		expect      *ListNode
	}{
		{
			description: "Test Case 1",
			input:       array2LinkList([]int{1, 2, 3, 4, 5}),
			expect:      array2LinkList([]int{5, 4, 3, 2, 1}),
		},
		{
			description: "Test Case 2",
			input:       array2LinkList([]int{1}),
			expect:      array2LinkList([]int{1}),
		},
		{
			description: "Test Case 1",
			input:       array2LinkList([]int{}),
			expect:      array2LinkList([]int{}),
		},
	}

	for _, t := range testCases {
		result := reverseLinkList(t.input)
		if !reflect.DeepEqual(result, t.expect) {
			log.Fatalf(
				"%s: [expect] != [actual]",
				t.description,
			)
		}
	}
}
