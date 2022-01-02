package main

import "mic/leetcode_100"

type TreeNode = leetcode_100.TreeNode


func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func( *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		node.Val += sum
		sum = node.Val
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func main()  {
	
}
