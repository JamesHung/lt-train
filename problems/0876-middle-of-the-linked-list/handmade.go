package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode is intentionally left blank for practice.
func middleNode(head *ListNode) *ListNode {
	oneStep, twoStep := head, head

	for twoStep != nil && twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}
	return oneStep
}
