type ListNode struct {
	val int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		val: val,
		next: next,
	}
}

func NewLinkedList() *LinkedList {
	empty := NewListNode(-1, nil)
	return &LinkedList{
		head: empty,
		tail: empty,
	}
}

func (ll *LinkedList) Get(index int) int {
	cur := ll.head.next
	i := 0
	for cur != nil {
		if i == index {
			return cur.val
		}
		i++
		cur = cur.next
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newLn := NewListNode(val, ll.head.next)
	ll.head.next = newLn
	if newLn.next == nil {
		ll.tail = newLn
	}
}

func (ll *LinkedList) InsertTail(val int) {
	newLn := NewListNode(val, nil)
	ll.tail.next = newLn
	ll.tail = newLn
}

func (ll *LinkedList) Remove(index int) bool {
	cur := ll.head
	for i := 0; i < index; i++ {
		if cur.next == nil {
			return false
		}
		cur = cur.next
	}

	if cur != nil && cur.next != nil {
		if cur.next == ll.tail {
			ll.tail = cur
		}
		cur.next = cur.next.next
		return true
	}

	return false
}

func (ll *LinkedList) GetValues() []int {
	res := []int{}
	cur := ll.head.next

	for cur != nil {
		res = append(res, cur.val)
		cur = cur.next
	}

	return res
}
