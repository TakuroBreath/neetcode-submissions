/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

    cur, next := head, head.Next
    head.Next = nil

	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}

	return cur
}
