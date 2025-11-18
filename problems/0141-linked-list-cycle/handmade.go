package linkedlistcycle

// ListNode defines a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle detects a cycle using Floyd's tortoise and hare algorithm.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false
}

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	current := head
// 	visited := make(map[*ListNode]bool)
// 	for current != nil {
// 		if visited[current] {
// 			return true
// 		}
// 		visited[current] = true
// 		current = current.Next
// 	}
// 	return false
// }
