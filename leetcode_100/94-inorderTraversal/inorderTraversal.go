package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var InorderTraversal func(*TreeNode)
	InorderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		InorderTraversal(node.Left)
		res = append(res, node.Val)
		InorderTraversal(node.Right)
	}
	InorderTraversal(root)
	return res
}

func main() {

}
