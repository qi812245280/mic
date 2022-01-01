package main

import "mic/leetcode_100"

type ListNode = leetcode_100.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	res := &ListNode{}
	res.Next = list1
	cur := res
	for list2 != nil {
		node1 := cur.Next
		if node1 == nil {
			cur.Next = list2
			break
		}
		if list2.Val < node1.Val {
			next2 := list2.Next
			cur.Next = list2
			list2.Next = node1
			list2 = next2
		}
		cur = cur.Next
	}
	return res.Next
}

func main() {
	node1_1 := &ListNode{Val: 1}
	node1_2 := &ListNode{Val: 2}
	node1_3 := &ListNode{Val: 4}
	node1_1.Next = node1_2
	node1_2.Next = node1_3

	node2_1 := &ListNode{Val: 1}
	node2_2 := &ListNode{Val: 3}
	node2_3 := &ListNode{Val: 4}
	node2_1.Next = node2_2
	node2_2.Next = node2_3

	mergeTwoLists(node1_1, node2_1).Print()
}
