package main

import (
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	// 因為有可能兩點間最大距離有可能出現在子樹中()
	max := math.Max(float64(diameterOfBinaryTree(root.Left)), float64(diameterOfBinaryTree(root.Right)))

	return int(math.Max(float64(left+right), max))
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	return int(math.Max(float64(left), float64(right))) + 1
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
				Val: 1,
			},
			expect: 0,
		},
		{
			description: "Test Case 2",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 8,
							},
						},
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 8,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expect: 6,
		},
	}

	for _, t := range testCases {
		result := diameterOfBinaryTree(t.input)
		if result != t.expect {
			log.Fatalf(
				"%s: expect:[%v] != actual:[%v]",
				t.description, t.expect, result,
			)
		}
	}
}
