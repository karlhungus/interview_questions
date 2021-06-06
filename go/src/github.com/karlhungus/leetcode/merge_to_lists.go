package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pl1 := l1
	pl2 := l2
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	var cur *ListNode
	if pl1.Val >= pl2.Val {
		res = pl2
		pl2 = pl2.Next
	} else {
		res = pl1
		pl1 = pl1.Next
	}
	cur = res

	for {
		if pl1 == nil {
			cur.Next = pl2
			return res
		}
		if pl2 == nil {
			cur.Next = pl1
			return res
		}
		if pl1.Val >= pl2.Val {
			cur.Next = pl2
			pl2 = pl2.Next
		} else {
			cur.Next = pl1
			pl1 = pl1.Next
		}
		cur = cur.Next
	}
}
