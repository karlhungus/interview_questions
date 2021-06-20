package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nhead := head.Next
	first := head
	second := first.Next
	var last *ListNode
	for first != nil && first.Next != nil {
		first.Next = second.Next
		second.Next = first
		if last != nil {
			last.Next = second
		}

		if first.Next != nil {
			last = first
			first = first.Next
			second = first.Next

		}
	}
	return nhead
}
