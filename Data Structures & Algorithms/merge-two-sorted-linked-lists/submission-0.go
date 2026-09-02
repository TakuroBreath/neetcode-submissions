/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1 
	}

	node := &ListNode{}
	res := node
	
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
        res = res.Next
	}

    if list1 != nil {
        res.Next = list1
    } else {
        res.Next = list2
    }

	return node.Next
}
