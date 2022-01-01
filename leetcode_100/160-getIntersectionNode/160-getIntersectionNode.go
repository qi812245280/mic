package main

import (
	"fmt"
	"mic/leetcode_100"
)

type ListNode = leetcode_100.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var lenA, lenB int
	var nodeA, nodeB = headA, headB
	for nodeA != nil {
		lenA += 1
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		lenB += 1
		nodeB = nodeB.Next
	}
	nodeA, nodeB = headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			nodeA = nodeA.Next
		}
	}
	if lenB > lenA {
		for i := 0; i < lenB-lenA; i++ {
			nodeB = nodeB.Next
		}
	}
	for nodeA != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nil
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}

	node5.Next = node6
	node6.Next = node7
	node1.Next = node5
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	res := getIntersectionNode(node1, node2)
	if res != nil {
		fmt.Println(res.Val)
	} else {
		fmt.Println("不存在")
	}
}
