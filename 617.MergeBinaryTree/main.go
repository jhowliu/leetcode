package main

import (
	"log"
	"reflect"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func merge(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	node := new(TreeNode)

	if t1 != nil && t2 != nil {
		node.Val = t1.Val + t2.Val
		node.Right = merge(t1.Right, t2.Right)
		node.Left = merge(t1.Left, t2.Left)
	} else if t1 != nil {
		node.Val = t1.Val
		node.Right = merge(t1.Right, nil)
		node.Left = merge(t1.Left, nil)
	} else {
		node.Val = t2.Val
		node.Right = merge(t2.Right, nil)
		node.Left = merge(t2.Left, nil)
	}

	return node
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	result := new(TreeNode)

	result = merge(t1, t2)

	return result
}

func main() {
	testCases := []struct {
		description string
		t1          *TreeNode
		t2          *TreeNode
		expect      *TreeNode
	}{
		{
			description: "testing 1",
			t1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			t2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expect: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		actual := mergeTrees(tt.t1, tt.t2)
		if !reflect.DeepEqual(tt.expect, actual) {
			log.Fatalf("%s: expect:%v != actual:%v", tt.description, tt.expect, actual)
		}
	}
}
