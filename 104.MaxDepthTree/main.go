package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	testCases := []struct {
		description string
		input       *TreeNode
		expect      int
	}{
		{
			description: "Test Case 1",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expect: 3,
		},
	}

	for _, t := range testCases {
		result := maxDepth(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: input[%v], expect[%v] != actual[%v]",
				t.description, t.input, t.expect, result,
			)
		}
	}
}
