package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func InvertTree2(root *TreeNode) {
	if root == nil {
		return
	}
	nodeList := []*TreeNode{root}
	for len(nodeList) > 0 {
		tmpList := make([]*TreeNode, 0)
		for _, node := range nodeList {
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				tmpList = append(tmpList, node.Left)
			}
			if node.Right != nil {
				tmpList = append(tmpList, node.Right)
			}
		}
		nodeList = tmpList
	}
}

// 递归
func InvertTree(node *TreeNode) {
	if node == nil {
		return
	}
	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp
	InvertTree(node.Left)
	InvertTree(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	// 1.递归
	InvertTree(root)
	return root
}

func main() {
	//root := TreeNode{Val: 4}
	//node1 := TreeNode{Val: 2}
}
