package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}
	var cur, pos *ListNode
	for cur = head.Next; cur.Next != nil; cur = cur.Next {
		if cur.Next == cur {
			return true
		}
		for pos = head; pos != cur; pos = pos.Next {
			if cur.Next == pos {
				return true
			}
		}

	}
	return false
}
