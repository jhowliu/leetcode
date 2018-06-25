package main

import (
	"log"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(tmp)

	return root
}

func main() {
	testCases := []struct {
		description string
		input       *TreeNode
		expect      *TreeNode
	}{
		{
			description: "testing 1",
			input: &TreeNode{
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
			expect: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, t := range testCases {
		result := invertTree(t.input)
		if !reflect.DeepEqual(t.expect, result) {
			log.Fatalf("%s: expect:%v != actual:%v", t.description, t.expect, result)
		}
	}
}
