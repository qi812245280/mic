package main

import (
	"mic/leetcode_100"
)

type ListNode = leetcode_100.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
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

var res *ListNode

func reverseList2(head *ListNode) *ListNode {
	return dfs(nil, head)
}

func dfs(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return dfs(cur, next)
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	reverseList2(l1).Print()
}
