package main

import (
	"fmt"
	"math"
	"mic/leetcode_100"
)

type TreeNode = leetcode_100.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode, *TreeNode, []*TreeNode)
	res := make([]*TreeNode, 0)
	dfs = func(node *TreeNode, target *TreeNode, path []*TreeNode) {
		if node == target{
			path = append(path, node)
			res = make([]*TreeNode, len(path))
			copy(res, path)
			return
		}
		path  = append(path, node)
		if node.Left != nil{
			dfs(node.Left, target, path)
			if len(res) != 0 {
				return
			}
		}
		if node.Right!= nil {
			dfs(node.Right, target, path)
			if len(res) != 0 {
				return
			}
		}
	}
	dfs(root, p, []*TreeNode{})
	pRes := make([]*TreeNode, len(res))
	copy(pRes, res)
	res = make([]*TreeNode, 0)

	dfs(root, q, []*TreeNode{})
	qRes := make([]*TreeNode, len(res))
	copy(qRes, res)

	var resNode  *TreeNode
	length := int(math.Min(float64(len(pRes)), float64(len(qRes))))
	for i := 0; i <length;i ++ {
		if pRes[i] == qRes[i] {
			resNode = pRes[i]
		}
	}
	return resNode
}


func main()  {
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node3.Left = node5
	node3.Right = node1
	node5.Left = node6
	node5.Right = node2
	node2.Left = node7
	node2.Right = node4
	node1.Left=node0
	node1.Right = node8
	res := lowestCommonAncestor(node3, node5, node1)
	fmt.Println(res.Val)
}

