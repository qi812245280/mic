package main

import (
	"fmt"
	"mic/leetcode_100"
)

type ListNode = leetcode_100.ListNode

func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	tmp1 := head
	var length int
	for tmp1 != nil {
		length++
		tmp1 = tmp1.Next
	}
	pathLen := length / 2
	fmt.Println(length)
	node2 := head
	for pathLen > 0 {
		node2 = node2.Next
		pathLen--
	}
	node3 := reverseListNode(node2)
	head.Print()
	node3.Print()
	for head != nil && node3 != nil {
		if head.Val != node3.Val {
			return false
		}
		head = head.Next
		node3 = node3.Next
	}
	return true
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 1}
	//node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node5
	isPalindrome(node1)
}
