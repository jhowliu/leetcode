package main

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	t := map[int]int{}

	if root == nil {
		return false
	}

	return dfs(root, t, k)
}

func dfs(root *TreeNode, t map[int]int, k int) bool {
	if root == nil {
		return false
	}

	if _, ok := t[root.Val]; ok {
		return true
	} else {
		t[k-root.Val] = root.Val
	}

	return dfs(root.Left, t, k) || dfs(root.Right, t, k)
}

func main() {
	testCases := []struct {
		description string
		input       *TreeNode
		target      int
		expect      bool
	}{
		{
			description: "testing 1",
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			target: 9,
			expect: true,
		},
	}

	for _, t := range testCases {
		if result := findTarget(t.input, t.target); t.expect != result {
			log.Fatalf("%s: expect:%v != actual:%v", t.description, t.expect, result)
		}
	}
}
