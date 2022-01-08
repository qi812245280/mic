package main

import "mic/leetcode_100"

type ListNode = leetcode_100.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next == nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	r := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(r)
	return merge(left, right)
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			cur.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			cur.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val <= list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
		}
		cur = cur.Next
	}
	return res.Next
}

func main() {

}
