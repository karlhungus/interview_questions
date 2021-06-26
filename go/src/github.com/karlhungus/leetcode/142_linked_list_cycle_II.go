package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	counter := map[*ListNode]struct{}{}
	cur := head
	if cur == nil || cur.Next == cur {
		return cur
	}
	for {
		if cur == nil {
			return nil
		}
		if _, ok := counter[cur]; ok {
			return cur
		} else {
			counter[cur] = struct{}{}
		}
		cur = cur.Next
	}
}
