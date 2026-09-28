/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

    fast, slow := head.Next, head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		fast, slow = fast.Next.Next, slow.Next
	}

	return false
}
