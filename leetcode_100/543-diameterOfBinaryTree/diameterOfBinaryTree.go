package main

import (
	"mic/leetcode_100"
	"math"
)

type TreeNode = leetcode_100.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxValue := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		maxValue = int(math.Max(float64(maxValue), float64(l+r)))
		return 1 + int(math.Max(float64(l), float64(r)))
	}
	dfs(root)
	return maxValue
}

func main() {

}
