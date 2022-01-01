package main

import (
	"mic/leetcode_100"
	"math"
)

type TreeNode = leetcode_100.TreeNode

func help(node1, node2 *TreeNode) bool {
	if node1 == nil || node2 == nil {
		if node1 == nil && node2 == nil {
			return true
		}
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return help(node1.Left, node2.Right) && help(node1.Right, node2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return help(root.Left, root.Right)
}

// TODO
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	idx := 0
	Nodes := []*TreeNode{root}
	for len(Nodes) > 0 {
		length := int(math.Pow(2, float64(idx)))
		if len(Nodes) != length {
			return false
		}
		for i := 0; i < len(Nodes)/2; i++ {
			if Nodes[i].Val != Nodes[length-1-i].Val {
				return false
			}
		}
		tmp := make([]*TreeNode, 0)
		for _, node := range Nodes {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		Nodes = tmp
		idx++
	}
	return true
}

func main() {

}
