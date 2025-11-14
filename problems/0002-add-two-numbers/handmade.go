package addtwonumbers

// ListNode is a singly-linked list node used by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers sums the values represented by l1 and l2 and returns the result as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	currentNode := dummyNode
	carryNum := 0
	for l1 != nil || l2 != nil || carryNum != 0 {
		sum := carryNum
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carryNum = sum / 10
		newNode := &ListNode{Val: sum % 10}
		currentNode.Next = newNode
		currentNode = newNode
	}

	return dummyNode.Next
}
