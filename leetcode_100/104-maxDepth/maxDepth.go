package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deep := 0
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		tmp := make([]*TreeNode, 0)
		for _, node := range nodeList {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nodeList = tmp
		deep += 1
	}
	return deep
}

func main() {

}
