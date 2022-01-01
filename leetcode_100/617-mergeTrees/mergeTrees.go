package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MergeTrees(node1, node2 *TreeNode) {
	node1.Val += node2.Val
	if node2.Left != nil {
		if node1.Left == nil {
			node1.Left = &TreeNode{}
		}
		MergeTrees(node1.Left, node2.Left)
	}
	if node2.Right != nil {
		if node1.Right == nil {
			node1.Right = &TreeNode{}
		}
		MergeTrees(node1.Right, node2.Right)
	}
}

func MergeTrees2(node1, node2 *TreeNode) *TreeNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	node1.Val += node2.Val
	node1.Left = MergeTrees2(node1.Left, node2.Left)
	node1.Right = MergeTrees2(node1.Right, node2.Right)
	return node1
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	MergeTrees(root1, root2)
	return root1
}

func main() {

}
