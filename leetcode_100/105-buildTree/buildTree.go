package main

import "mic/leetcode_100"

type TreeNode = leetcode_100.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	idx := 0
	for i, num := range inorder {
		if num == node.Val {
			idx = i
			break
		}
	}
	lLen := len(inorder[:idx])
	node.Left = buildTree(preorder[1:1+lLen], inorder[:idx])
	node.Right = buildTree(preorder[1+lLen:], inorder[idx+1:])
	return node
}

func main() {

}
