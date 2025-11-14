package reverselinkedlist

const debugReverseLinkedList = false

// ListNode defines a singly linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList reverses a list iteratively and returns the new head.
func reverse(head *ListNode) *ListNode {
	return reverseListRecursive(head)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
