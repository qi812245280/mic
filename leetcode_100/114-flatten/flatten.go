package main

import (
	"mic/leetcode_100"
)

type TreeNode = leetcode_100.TreeNode

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	flatten(root.Left)
	rightBk := root.Right
	if root.Left != nil {
		// 左边做好了，把左边的树插入到root和root.Right之间
		l := root.Left
		for l.Right != nil {
			l = l.Right
		}
		l.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	flatten(rightBk)
}


func main()  {
	
}
